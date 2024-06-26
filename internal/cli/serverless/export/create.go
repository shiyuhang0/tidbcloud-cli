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
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/juju/errors"
	"github.com/spf13/cobra"

	"tidbcloud-cli/internal"
	"tidbcloud-cli/internal/config"
	"tidbcloud-cli/internal/flag"
	"tidbcloud-cli/internal/service/cloud"
	"tidbcloud-cli/internal/ui"
	"tidbcloud-cli/internal/util"
	exportApi "tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless_export/client/export_service"
	exportModel "tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless_export/models"
)

type TargetType string

const (
	TargetTypeS3      TargetType = "S3"
	TargetTypeLOCAL   TargetType = "LOCAL"
	TargetTypeUnknown TargetType = "UNKNOWN"
)

type FileType string

const (
	FileTypeSQL     FileType = "SQL"
	FileTypeCSV     FileType = "CSV"
	FileTypeUnknown FileType = "UNKNOWN"
)

var S3InputFields = map[string]int{
	flag.S3URI:             0,
	flag.S3AccessKeyID:     1,
	flag.S3SecretAccessKey: 2,
}

var FilterSQLInputFields = map[string]int{
	flag.SQL: 0,
}

var FilterTableInputFields = map[string]int{
	flag.TableFilter: 0,
	flag.TableWhere:  1,
}

type CreateOpts struct {
	interactive bool
}

func (c CreateOpts) NonInteractiveFlags() []string {
	return []string{
		flag.ClusterID,
		flag.FileType,
		flag.TargetType,
		flag.S3URI,
		flag.S3AccessKeyID,
		flag.S3SecretAccessKey,
		flag.Compression,
		flag.SQL,
		flag.TableFilter,
		flag.TableWhere,
	}
}

func (c CreateOpts) RequiredFlags() []string {
	return []string{
		flag.ClusterID,
	}
}

func (c *CreateOpts) MarkInteractive(cmd *cobra.Command) error {
	flags := c.NonInteractiveFlags()
	for _, fn := range flags {
		f := cmd.Flags().Lookup(fn)
		if f != nil && f.Changed {
			c.interactive = false
			break
		}
	}
	// Mark required flags
	if !c.interactive {
		for _, fn := range c.RequiredFlags() {
			err := cmd.MarkFlagRequired(fn)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CreateCmd(h *internal.Helper) *cobra.Command {
	var force bool
	opts := CreateOpts{
		interactive: true,
	}

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Export data from a TiDB Serverless cluster",
		Args:  cobra.NoArgs,
		Example: fmt.Sprintf(`  Create an export in interactive mode:
  $ %[1]s serverless export create

  Export all data with local type in non-interactive mode:
  $ %[1]s serverless export create -c <cluster-id>

  Export all data with s3 type in non-interactive mode:
  $ %[1]s serverless export create -c <cluster-id> --s3.uri <s3-uri> --s3.access-key-id <access-key-id> --s3.secret-access-key <secret-access-key>

  Export test.t1 and test.t2 in non-interactive mode:
  $ %[1]s serverless export create -c <cluster-id> --filter 'test.t1,test.t2'

  Export tables with special characters, for example, if you want to export %[2]stest,%[2]s.%[2]st1%[2]s and %[2]s"test%[2]s.%[2]st1%[2]s:
  $ %[1]s serverless export create -c <cluster-id> --filter '"%[2]stest1,%[2]s.t1","%[2]s""test%[2]s.t1"'`,
			config.CliName, "`"),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.MarkInteractive(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			d, err := h.Client()
			if err != nil {
				return err
			}
			ctx := cmd.Context()

			var s3URI, accessKeyID, secretAccessKey, targetType, fileType, compression string
			var clusterId, sql, where string
			var patterns []string
			if opts.interactive {
				if !h.IOStreams.CanPrompt {
					return errors.New("The terminal doesn't support interactive mode, please use non-interactive mode")
				}

				project, err := cloud.GetSelectedProject(ctx, h.QueryPageSize, d)
				if err != nil {
					return err
				}

				cluster, err := cloud.GetSelectedCluster(ctx, project.ID, h.QueryPageSize, d)
				if err != nil {
					return err
				}
				clusterId = cluster.ID

				selectedTargetType, err := GetSelectedTargetType()
				if err != nil {
					return err
				}
				if selectedTargetType == TargetTypeUnknown {
					return errors.New("target type must be LOCAL or S3")
				}
				targetType = string(selectedTargetType)

				if selectedTargetType == TargetTypeS3 {
					s3InputModel, err := GetS3Input()
					if err != nil {
						return err
					}
					s3URI = s3InputModel.(ui.TextInputModel).Inputs[S3InputFields[flag.S3URI]].Value()
					if s3URI == "" {
						return errors.New("s3 uri is required when target type is S3")
					}
					accessKeyID = s3InputModel.(ui.TextInputModel).Inputs[S3InputFields[flag.S3AccessKeyID]].Value()
					if accessKeyID == "" {
						return errors.New("access key Id is required when target type is S3")
					}
					secretAccessKey = s3InputModel.(ui.TextInputModel).Inputs[S3InputFields[flag.S3SecretAccessKey]].Value()
					if secretAccessKey == "" {
						return errors.New("secret access key is required when target type is S3")
					}
				}

				filterType, err := GetSelectedFilterType()
				if err != nil {
					return err
				}
				switch filterType {
				case FilterNone:
				case FilterSQL:
					filterInputModel, err := GetFilterInput(FilterSQL)
					if err != nil {
						return err
					}
					sql = filterInputModel.(ui.TextInputModel).Inputs[FilterSQLInputFields[flag.SQL]].Value()
					if sql == "" {
						return errors.New("sql is empty")
					}
				case FilterTable:
					fmt.Fprintln(h.IOStreams.Out, color.BlueString("Please input the following options, require at least one field"))
					filterInputModel, err := GetFilterInput(FilterTable)
					if err != nil {
						return err
					}
					// TODO input slice
					patternString := filterInputModel.(ui.TextInputModel).Inputs[FilterTableInputFields[flag.TableFilter]].Value()
					patterns, err = util.StringSliceConv(patternString)
					if err != nil {
						return err
					}
					where = filterInputModel.(ui.TextInputModel).Inputs[FilterTableInputFields[flag.TableWhere]].Value()
					if len(patterns) == 0 && where == "" {
						return errors.New("both patterns and where are empty, require at least one field")
					}
				}

				if filterType == FilterSQL {
					fileType = string(FileTypeCSV)
				} else {
					selectedFileType, err := GetSelectedFileType()
					if err != nil {
						return err
					}
					if selectedFileType == FileTypeUnknown {
						return errors.New("file type must be LOCAL or S3")
					}
					fileType = string(selectedFileType)
				}

				compression, err = GetSelectedCompression()
				if err != nil {
					return err
				}
			} else {
				// non-interactive mode, get values from flags
				var err error
				clusterId, err = cmd.Flags().GetString(flag.ClusterID)
				if err != nil {
					return errors.Trace(err)
				}
				targetType, err = cmd.Flags().GetString(flag.TargetType)
				if err != nil {
					return errors.Trace(err)
				}
				fileType, err = cmd.Flags().GetString(flag.FileType)
				if err != nil {
					return errors.Trace(err)
				}
				if targetType == string(TargetTypeS3) {
					s3URI, err = cmd.Flags().GetString(flag.S3URI)
					if err != nil {
						return errors.Trace(err)
					}
					if s3URI == "" {
						return errors.New("s3 uri is required when target type is S3")
					}
					accessKeyID, err = cmd.Flags().GetString(flag.S3AccessKeyID)
					if err != nil {
						return errors.Trace(err)
					}
					if accessKeyID == "" {
						return errors.New("accessKeyId is required when target type is S3")
					}
					secretAccessKey, err = cmd.Flags().GetString(flag.S3SecretAccessKey)
					if err != nil {
						return errors.Trace(err)
					}
					if secretAccessKey == "" {
						return errors.New("secretAccessKey is required when target type is S3")
					}
				}
				compression, err = cmd.Flags().GetString(flag.Compression)
				if err != nil {
					return errors.Trace(err)
				}
				sql, err = cmd.Flags().GetString(flag.SQL)
				if err != nil {
					return errors.Trace(err)
				}
				where, err = cmd.Flags().GetString(flag.TableWhere)
				if err != nil {
					return errors.Trace(err)
				}
				patterns, err = cmd.Flags().GetStringSlice(flag.TableFilter)
				if err != nil {
					return errors.Trace(err)
				}
			}

			if !opts.interactive && sql == "" && len(patterns) == 0 && !force {
				if !h.IOStreams.CanPrompt {
					return fmt.Errorf("the terminal doesn't support prompt, please run with --force to create export")
				}

				confirmationMessage := fmt.Sprintf("%s %s %s %s", color.BlueString("You will export the whole cluster."), color.BlueString("Please type"), color.HiBlueString(confirmed), color.BlueString("to continue:"))
				prompt := &survey.Input{
					Message: confirmationMessage,
				}
				var userInput string
				err := survey.AskOne(prompt, &userInput)
				if err != nil {
					if err == terminal.InterruptErr {
						return util.InterruptError
					} else {
						return err
					}
				}
				if userInput != confirmed {
					return errors.New("incorrect confirm string entered, skipping create")
				}
			}

			params := exportApi.NewExportServiceCreateExportParams().WithClusterID(clusterId).WithBody(
				exportApi.ExportServiceCreateExportBody{
					ExportOptions: &exportModel.V1beta1ExportOptions{
						FileType: exportModel.V1beta1ExportOptionsFileType(fileType),
					},
					Target: &exportModel.V1beta1Target{
						Type: exportModel.TargetTargetType(targetType),
						S3: &exportModel.TargetS3Target{
							URI: s3URI,
							AccessKey: &exportModel.S3TargetAccessKey{
								ID:     accessKeyID,
								Secret: secretAccessKey,
							},
						},
					},
				}).WithContext(ctx)
			if compression != "" {
				params.Body.ExportOptions.Compression = exportModel.ExportOptionsCompressionType(compression)
			}
			if sql != "" {
				params.Body.ExportOptions.Filter = &exportModel.ExportOptionsFilter{
					SQL: sql,
				}
			}
			if len(patterns) > 0 || where != "" {
				params.Body.ExportOptions.Filter = &exportModel.ExportOptionsFilter{
					Table: &exportModel.FilterTable{
						Where:    where,
						Patterns: patterns,
					},
				}
			}
			resp, err := d.CreateExport(params)
			if err != nil {
				return errors.Trace(err)
			}
			_, err = fmt.Fprintln(h.IOStreams.Out, color.GreenString("export %s is running now", resp.Payload.ExportID))
			if err != nil {
				return err
			}
			return nil
		},
	}

	createCmd.Flags().StringP(flag.ClusterID, flag.ClusterIDShort, "", "The ID of the cluster, in which the export will be created.")
	createCmd.Flags().String(flag.FileType, "CSV", "The export file type. One of [\"CSV\" \"SQL\"].")
	createCmd.Flags().String(flag.TargetType, "LOCAL", "The export target. One of [\"LOCAL\" \"S3\"].")
	createCmd.Flags().String(flag.S3URI, "", "The s3 uri in s3://<bucket>/<path> format. Required when target type is S3.")
	createCmd.Flags().String(flag.S3AccessKeyID, "", "The access key ID of the S3. Required when target type is S3.")
	createCmd.Flags().String(flag.S3SecretAccessKey, "", "The secret access key of the S3. Required when target type is S3.")
	createCmd.Flags().String(flag.Compression, "GZIP", "The compression algorithm of the export file. One of [\"GZIP\" \"SNAPPY\" \"ZSTD\" \"NONE\"].")
	createCmd.Flags().StringSlice(flag.TableFilter, nil, "Specify the exported table(s) with table filter patterns. See https://docs.pingcap.com/tidb/stable/table-filter to learn table filter.")
	createCmd.Flags().String(flag.TableWhere, "", "Filter the exported table(s) with the where condition.")
	createCmd.Flags().String(flag.SQL, "", "Filter the exported data with SQL SELECT statement.")
	createCmd.Flags().BoolVar(&force, flag.Force, false, "Create without confirmation. You need to confirm when you want to export the whole cluster.")
	createCmd.MarkFlagsMutuallyExclusive(flag.TableFilter, flag.SQL)
	createCmd.MarkFlagsMutuallyExclusive(flag.TableWhere, flag.SQL)
	return createCmd
}

func GetSelectedTargetType() (TargetType, error) {
	targetTypes := make([]interface{}, 0, 2)
	targetTypes = append(targetTypes, TargetTypeLOCAL, TargetTypeS3)
	model, err := ui.InitialSelectModel(targetTypes, "Choose export target:")
	if err != nil {
		return TargetTypeUnknown, errors.Trace(err)
	}

	p := tea.NewProgram(model)
	targetTypeModel, err := p.Run()
	if err != nil {
		return TargetTypeUnknown, errors.Trace(err)
	}
	if m, _ := targetTypeModel.(ui.SelectModel); m.Interrupted {
		return TargetTypeUnknown, util.InterruptError
	}
	targetType := targetTypeModel.(ui.SelectModel).GetSelectedItem()
	if targetType == nil {
		return TargetTypeUnknown, errors.New("no export target selected")
	}
	return targetType.(TargetType), nil
}

func GetSelectedFileType() (FileType, error) {
	fileTypes := make([]interface{}, 0, 2)
	fileTypes = append(fileTypes, FileTypeSQL, FileTypeCSV)
	model, err := ui.InitialSelectModel(fileTypes, "Choose export file type:")
	if err != nil {
		return FileTypeUnknown, errors.Trace(err)
	}

	p := tea.NewProgram(model)
	fileTypeModel, err := p.Run()
	if err != nil {
		return FileTypeUnknown, errors.Trace(err)
	}
	if m, _ := fileTypeModel.(ui.SelectModel); m.Interrupted {
		return FileTypeUnknown, util.InterruptError
	}
	fileType := fileTypeModel.(ui.SelectModel).GetSelectedItem()
	if fileType == nil {
		return FileTypeUnknown, errors.New("no export file type selected")
	}
	return fileType.(FileType), nil
}

func GetSelectedCompression() (string, error) {
	compressions := make([]interface{}, 0, 4)
	compressions = append(compressions, "GZIP", "SNAPPY", "ZSTD", "NONE")
	model, err := ui.InitialSelectModel(compressions, "Choose the compression algorithm:")
	if err != nil {
		return "", errors.Trace(err)
	}

	p := tea.NewProgram(model)
	fileTypeModel, err := p.Run()
	if err != nil {
		return "", errors.Trace(err)
	}
	if m, _ := fileTypeModel.(ui.SelectModel); m.Interrupted {
		return "", util.InterruptError
	}
	compression := fileTypeModel.(ui.SelectModel).GetSelectedItem()
	if compression == nil {
		return "", errors.New("no compression algorithm selected")
	}
	return compression.(string), nil
}

type FilterType string

const (
	FilterNone  FilterType = "Export all data"
	FilterTable FilterType = "Table (export specific database(s) or table(s) using table filter)"
	FilterSQL   FilterType = "SQL (export data using SELECT SQL statement)"
)

func GetSelectedFilterType() (FilterType, error) {
	filterTypes := make([]interface{}, 0, 3)

	filterTypes = append(filterTypes, FilterTable, FilterSQL, FilterNone)
	model, err := ui.InitialSelectModel(filterTypes, "Choose the filter type")
	if err != nil {
		return "", errors.Trace(err)
	}

	p := tea.NewProgram(model)
	fileTypeModel, err := p.Run()
	if err != nil {
		return "", errors.Trace(err)
	}
	if m, _ := fileTypeModel.(ui.SelectModel); m.Interrupted {
		return "", util.InterruptError
	}
	filterType := fileTypeModel.(ui.SelectModel).GetSelectedItem()
	if filterType == nil {
		return "", errors.New("no filter type selected")
	}
	return filterType.(FilterType), nil
}

func initialS3InputModel() ui.TextInputModel {
	m := ui.TextInputModel{
		Inputs: make([]textinput.Model, len(S3InputFields)),
	}
	for k, v := range S3InputFields {
		t := textinput.New()
		switch k {
		case flag.S3URI:
			t.Placeholder = "S3 URI in s3://<bucket>/<path> format"
			t.Focus()
			t.PromptStyle = config.FocusedStyle
			t.TextStyle = config.FocusedStyle
		case flag.S3AccessKeyID:
			t.Placeholder = "S3 Access Key ID"
		case flag.S3SecretAccessKey:
			t.Placeholder = "S3 Secret Access key"
		}
		m.Inputs[v] = t
	}
	return m
}

func GetS3Input() (tea.Model, error) {
	p := tea.NewProgram(initialS3InputModel())
	inputModel, err := p.Run()
	if err != nil {
		return nil, errors.Trace(err)
	}
	if inputModel.(ui.TextInputModel).Interrupted {
		return nil, util.InterruptError
	}
	return inputModel, nil
}

func initialFilterInputModel(filterType FilterType) ui.TextInputModel {
	var inputFields map[string]int
	switch filterType {
	case FilterTable:
		inputFields = FilterTableInputFields
	case FilterSQL:
		inputFields = FilterSQLInputFields
	}
	m := ui.TextInputModel{
		Inputs: make([]textinput.Model, len(inputFields)),
	}
	for k, v := range inputFields {
		t := textinput.New()
		switch k {
		case flag.SQL:
			t.Placeholder = "SELECT SQL statement"
			t.Focus()
			t.PromptStyle = config.FocusedStyle
			t.TextStyle = config.FocusedStyle
		case flag.TableFilter:
			t.Placeholder = "Table filter patterns (comma separated). Example: test1.t1,test.*"
			t.Focus()
			t.PromptStyle = config.FocusedStyle
			t.TextStyle = config.FocusedStyle
		case flag.TableWhere:
			t.Placeholder = "Where condition. Example: id > 10"
		}
		m.Inputs[v] = t
	}
	return m
}

func GetFilterInput(filterType FilterType) (tea.Model, error) {
	p := tea.NewProgram(initialFilterInputModel(filterType))
	inputModel, err := p.Run()
	if err != nil {
		return nil, errors.Trace(err)
	}
	if inputModel.(ui.TextInputModel).Interrupted {
		return nil, util.InterruptError
	}
	return inputModel, nil
}
