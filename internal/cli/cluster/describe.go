// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"encoding/json"
	"fmt"

	"tidbcloud-cli/internal"
	"tidbcloud-cli/internal/config"
	"tidbcloud-cli/internal/flag"
	"tidbcloud-cli/internal/service/cloud"
	"tidbcloud-cli/internal/telemetry"

	serverlessApi "tidbcloud-cli/pkg/tidbcloud/serverless/client/serverless_service"

	"github.com/juju/errors"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	interactive bool
}

func (c DescribeOpts) NonInteractiveFlags() []string {
	return []string{
		flag.ClusterID,
	}
}

func DescribeCmd(h *internal.Helper) *cobra.Command {
	opts := DescribeOpts{
		interactive: true,
	}

	var describeCmd = &cobra.Command{
		Use:         "describe",
		Short:       "Describe a cluster",
		Aliases:     []string{"get"},
		Annotations: make(map[string]string),
		Example: fmt.Sprintf(`  Get the cluster info in interactive mode:
 $ %[1]s cluster describe

 Get the cluster info in non-interactive mode:
 $ %[1]s cluster describe -c <cluster-id>`, config.CliName),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			flags := opts.NonInteractiveFlags()
			for _, fn := range flags {
				f := cmd.Flags().Lookup(fn)
				if f != nil && f.Changed {
					opts.interactive = false
				}
			}

			// mark required flags in non-interactive mode
			if !opts.interactive {
				for _, fn := range flags {
					err := cmd.MarkFlagRequired(fn)
					if err != nil {
						return errors.Trace(err)
					}
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			d, err := h.Client()
			if err != nil {
				return err
			}

			var clusterID string
			if opts.interactive {
				cmd.Annotations[telemetry.InteractiveMode] = "true"
				if !h.IOStreams.CanPrompt {
					return errors.New("The terminal doesn't support interactive mode, please use non-interactive mode")
				}

				// interactive mode
				project, err := cloud.GetSelectedProject(h.QueryPageSize, d)
				if err != nil {
					return err
				}
				projectID := project.ID

				cluster, err := cloud.GetSelectedCluster(projectID, h.QueryPageSize, d)
				if err != nil {
					return err
				}
				clusterID = cluster.ID
			} else {
				// non-interactive mode does not need projectID
				cID, err := cmd.Flags().GetString(flag.ClusterID)
				if err != nil {
					return errors.Trace(err)
				}
				clusterID = cID
			}

			view, err := cmd.Flags().GetString(flag.View)
			if err != nil {
				return errors.Trace(err)
			}

			params := serverlessApi.NewServerlessServiceGetClusterParams().WithClusterID(clusterID)
			if view == flag.BasicView {
				params.WithView(&view)
			} else if view != flag.FullView {
				return errors.Errorf("invalid view: %s", view)
			}

			cluster, err := d.GetCluster(params)
			if err != nil {
				return errors.Trace(err)
			}

			v, err := json.MarshalIndent(cluster.Payload, "", "  ")
			if err != nil {
				return errors.Trace(err)
			}

			fmt.Fprintln(h.IOStreams.Out, string(v))
			return nil
		},
	}

	describeCmd.Flags().StringP(flag.ClusterID, flag.ClusterIDShort, "", "The ID of the cluster")
	describeCmd.Flags().StringP(flag.View, flag.ViewShort, flag.FullView, "The view of cluster, One of [\"BASIC\" \"FULL\"]")
	return describeCmd
}
