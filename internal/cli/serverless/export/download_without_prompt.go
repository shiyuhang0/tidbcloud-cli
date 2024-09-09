// Copyright 2024 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package export

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"

	"tidbcloud-cli/internal"
	"tidbcloud-cli/internal/config"
	"tidbcloud-cli/internal/service/cloud"
	"tidbcloud-cli/internal/util"
)

var wg = sync.WaitGroup{}

const BatchSize = 20

// downloadPool download export files concurrently
type downloadPool struct {
	path        string
	concurrency int
	client      cloud.TiDBCloudClient
	// The size of the batch to request download url
	clusterID string
	exportID  string
	count     int
	h         *internal.Helper

	jobs    chan *downloadJob
	results chan *downloadResult

	// fileJobs is a list of jobs that will be set to the jobs channel
	fileJobs  []*downloadJob
	batchSize int
}

func NewDownloadPool(h *internal.Helper, path string,
	concurrency int, exportID, clusterID string, count int, client cloud.TiDBCloudClient) (*downloadPool, error) {
	if count <= 0 {
		return nil, errors.New("no files to download")
	}
	if concurrency <= 0 {
		concurrency = DefaultConcurrency
	}

	jobBufferSize := 2 * concurrency
	if jobBufferSize > count {
		jobBufferSize = count
	}
	jobs := make(chan *downloadJob, jobBufferSize)
	results := make(chan *downloadResult, count)

	return &downloadPool{
		h:           h,
		path:        path,
		concurrency: concurrency,
		client:      client,
		clusterID:   clusterID,
		exportID:    exportID,
		jobs:        jobs,
		results:     results,
		batchSize:   BatchSize,
		count:       count,
	}, nil
}

type downloadJob struct {
	fileName    string
	downloadUrl *string
	size        int
	err         error
}

type downloadResult struct {
	name   string
	err    error
	status JobStatus
}

func (r *downloadResult) GetResult() string {
	if r.status == Succeeded {
		return fmt.Sprintf("%s success", r.name)
	}
	if r.err == nil {
		return fmt.Sprintf("%s %s", r.name, r.status)
	}
	return fmt.Sprintf("%s %s: %s", r.name, r.status, r.err.Error())
}

func (d *downloadPool) Start() error {
	// create the path if not exist
	err := util.CreateFolder(d.path)
	if err != nil {
		return err
	}
	pathName := d.path
	if pathName == "" {
		pathName = "current folder"
	}
	fmt.Fprintf(d.h.IOStreams.Out, color.GreenString("start to download files to %s:\n", pathName))
	// start produce
	go d.produce()
	// start consumers:
	for i := 0; i < d.concurrency; i++ {
		wg.Add(1)
		go d.consume()
	}
	// wait for all consumers to finish
	wg.Wait()
	close(d.results)
	// summarize the download results
	succeededCount := 0
	failedCount := 0
	skippedCount := 0
	downloadResults := make([]*downloadResult, 0)
	for result := range d.results {
		switch result.status {
		case Succeeded:
			succeededCount++
		case Failed:
			failedCount++
		case Skipped:
			skippedCount++
		}
		downloadResults = append(downloadResults, result)
	}
	fmt.Fprintf(d.h.IOStreams.Out, generateDownloadSummary(succeededCount, skippedCount, failedCount))
	index := 0
	for _, f := range downloadResults {
		if f.status != Succeeded {
			index++
			fmt.Fprintf(d.h.IOStreams.Out, "%d.%s\n", index, f.GetResult())
		}
	}
	if failedCount > 0 {
		return fmt.Errorf("%d file(s) failed to download", failedCount)
	}
	return nil
}

func (d *downloadPool) produce() {
	pageSize := int32(d.batchSize)
	var pageToken *string
	ctx := context.Background()
	for i := 0; i < d.count; i++ {
		// request the next batch when the fileJobs are not enough
		if len(d.fileJobs) < i+1 {
			resp, err := d.client.ListExportFilesWithRetry(ctx, d.clusterID, d.exportID, &pageSize, pageToken, true)
			if err != nil {
				// skip the rest of files
				errCount := d.count - len(d.fileJobs)
				for j := 0; j < errCount; j++ {
					d.fileJobs = append(d.fileJobs, &downloadJob{fileName: "unknown", err: err})
				}
			} else {
				pageToken = resp.NextPageToken
				for _, file := range resp.Files {
					job := &downloadJob{
						fileName:    *file.Name,
						downloadUrl: file.Url,
					}
					d.fileJobs = append(d.fileJobs, job)
				}
			}
		}
		// produce job
		d.jobs <- d.fileJobs[i]
	}
	close(d.jobs)
}

func (d *downloadPool) consume() {
	defer wg.Done()
	for job := range d.jobs {
		func() {
			var err error
			defer func() {
				// set result to channel
				if err != nil {
					if strings.Contains(err.Error(), "file already exists") {
						fmt.Fprintf(d.h.IOStreams.Out, "download %s skipped: %s\n", job.fileName, err.Error())
						d.results <- &downloadResult{name: job.fileName, err: err, status: Skipped}
					} else {
						fmt.Fprintf(d.h.IOStreams.Out, "download %s failed: %s\n", job.fileName, err.Error())
						d.results <- &downloadResult{name: job.fileName, err: err, status: Failed}
					}
				} else {
					fmt.Fprintf(d.h.IOStreams.Out, "download %s succeeded | %s\n", job.fileName, humanize.IBytes(uint64(job.size)))
					d.results <- &downloadResult{name: job.fileName, err: nil, status: Succeeded}
				}
			}()

			// request the url
			if job.err != nil {
				err = job.err
				return
			}
			if job.downloadUrl == nil {
				err = errors.New("empty download url")
				return
			}
			resp, err := util.GetResponse(*job.downloadUrl, os.Getenv(config.DebugEnv) != "")
			if err != nil {
				return
			}
			job.size = int(resp.ContentLength)
			defer resp.Body.Close()

			file, err := util.CreateFile(d.path, job.fileName)
			if err != nil {
				return
			}
			defer file.Close()
			_, err = io.Copy(file, resp.Body)
		}()
	}
}

func generateDownloadSummary(succeededCount, skippedCount, failedCount int) string {
	summaryMessage := fmt.Sprintf("%s %s %s", color.BlueString("download summary:"), color.GreenString("succeeded: %d", succeededCount), color.GreenString("skipped: %d", skippedCount))
	if failedCount > 0 {
		summaryMessage += color.RedString(" failed: %d", failedCount)
	} else {
		summaryMessage += fmt.Sprintf(" failed: %d", failedCount)
	}
	return summaryMessage + "\n"
}