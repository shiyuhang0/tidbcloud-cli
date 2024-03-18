// Code generated by go-swagger; DO NOT EDIT.

package export_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless_export/models"
)

// ExportServiceDownloadExportReader is a Reader for the ExportServiceDownloadExport structure.
type ExportServiceDownloadExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportServiceDownloadExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportServiceDownloadExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportServiceDownloadExportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportServiceDownloadExportOK creates a ExportServiceDownloadExportOK with default headers values
func NewExportServiceDownloadExportOK() *ExportServiceDownloadExportOK {
	return &ExportServiceDownloadExportOK{}
}

/*
ExportServiceDownloadExportOK describes a response with status code 200, with default header values.

A successful response.
*/
type ExportServiceDownloadExportOK struct {
	Payload *models.V1beta1DownloadExportsResponse
}

// IsSuccess returns true when this export service download export o k response has a 2xx status code
func (o *ExportServiceDownloadExportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export service download export o k response has a 3xx status code
func (o *ExportServiceDownloadExportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export service download export o k response has a 4xx status code
func (o *ExportServiceDownloadExportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export service download export o k response has a 5xx status code
func (o *ExportServiceDownloadExportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export service download export o k response a status code equal to that given
func (o *ExportServiceDownloadExportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export service download export o k response
func (o *ExportServiceDownloadExportOK) Code() int {
	return 200
}

func (o *ExportServiceDownloadExportOK) Error() string {
	return fmt.Sprintf("[GET /v1beta1/clusters/{clusterId}/exports/{exportId}:download][%d] exportServiceDownloadExportOK  %+v", 200, o.Payload)
}

func (o *ExportServiceDownloadExportOK) String() string {
	return fmt.Sprintf("[GET /v1beta1/clusters/{clusterId}/exports/{exportId}:download][%d] exportServiceDownloadExportOK  %+v", 200, o.Payload)
}

func (o *ExportServiceDownloadExportOK) GetPayload() *models.V1beta1DownloadExportsResponse {
	return o.Payload
}

func (o *ExportServiceDownloadExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1beta1DownloadExportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportServiceDownloadExportDefault creates a ExportServiceDownloadExportDefault with default headers values
func NewExportServiceDownloadExportDefault(code int) *ExportServiceDownloadExportDefault {
	return &ExportServiceDownloadExportDefault{
		_statusCode: code,
	}
}

/*
ExportServiceDownloadExportDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ExportServiceDownloadExportDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this export service download export default response has a 2xx status code
func (o *ExportServiceDownloadExportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export service download export default response has a 3xx status code
func (o *ExportServiceDownloadExportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export service download export default response has a 4xx status code
func (o *ExportServiceDownloadExportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export service download export default response has a 5xx status code
func (o *ExportServiceDownloadExportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export service download export default response a status code equal to that given
func (o *ExportServiceDownloadExportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export service download export default response
func (o *ExportServiceDownloadExportDefault) Code() int {
	return o._statusCode
}

func (o *ExportServiceDownloadExportDefault) Error() string {
	return fmt.Sprintf("[GET /v1beta1/clusters/{clusterId}/exports/{exportId}:download][%d] ExportService_DownloadExport default  %+v", o._statusCode, o.Payload)
}

func (o *ExportServiceDownloadExportDefault) String() string {
	return fmt.Sprintf("[GET /v1beta1/clusters/{clusterId}/exports/{exportId}:download][%d] ExportService_DownloadExport default  %+v", o._statusCode, o.Payload)
}

func (o *ExportServiceDownloadExportDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ExportServiceDownloadExportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}