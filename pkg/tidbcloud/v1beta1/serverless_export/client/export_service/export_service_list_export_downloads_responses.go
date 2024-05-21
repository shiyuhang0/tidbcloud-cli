// Code generated by go-swagger; DO NOT EDIT.

package export_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless_export/models"
)

// ExportServiceListExportDownloadsReader is a Reader for the ExportServiceListExportDownloads structure.
type ExportServiceListExportDownloadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportServiceListExportDownloadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportServiceListExportDownloadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportServiceListExportDownloadsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportServiceListExportDownloadsOK creates a ExportServiceListExportDownloadsOK with default headers values
func NewExportServiceListExportDownloadsOK() *ExportServiceListExportDownloadsOK {
	return &ExportServiceListExportDownloadsOK{}
}

/*
ExportServiceListExportDownloadsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ExportServiceListExportDownloadsOK struct {
	Payload *models.V1beta1ListExportDownloadsResponse
}

// IsSuccess returns true when this export service list export downloads o k response has a 2xx status code
func (o *ExportServiceListExportDownloadsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export service list export downloads o k response has a 3xx status code
func (o *ExportServiceListExportDownloadsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export service list export downloads o k response has a 4xx status code
func (o *ExportServiceListExportDownloadsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export service list export downloads o k response has a 5xx status code
func (o *ExportServiceListExportDownloadsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export service list export downloads o k response a status code equal to that given
func (o *ExportServiceListExportDownloadsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export service list export downloads o k response
func (o *ExportServiceListExportDownloadsOK) Code() int {
	return 200
}

func (o *ExportServiceListExportDownloadsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1beta1/clusters/{clusterId}/exports/{exportId}/downloads][%d] exportServiceListExportDownloadsOK %s", 200, payload)
}

func (o *ExportServiceListExportDownloadsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1beta1/clusters/{clusterId}/exports/{exportId}/downloads][%d] exportServiceListExportDownloadsOK %s", 200, payload)
}

func (o *ExportServiceListExportDownloadsOK) GetPayload() *models.V1beta1ListExportDownloadsResponse {
	return o.Payload
}

func (o *ExportServiceListExportDownloadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1beta1ListExportDownloadsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportServiceListExportDownloadsDefault creates a ExportServiceListExportDownloadsDefault with default headers values
func NewExportServiceListExportDownloadsDefault(code int) *ExportServiceListExportDownloadsDefault {
	return &ExportServiceListExportDownloadsDefault{
		_statusCode: code,
	}
}

/*
ExportServiceListExportDownloadsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ExportServiceListExportDownloadsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this export service list export downloads default response has a 2xx status code
func (o *ExportServiceListExportDownloadsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export service list export downloads default response has a 3xx status code
func (o *ExportServiceListExportDownloadsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export service list export downloads default response has a 4xx status code
func (o *ExportServiceListExportDownloadsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export service list export downloads default response has a 5xx status code
func (o *ExportServiceListExportDownloadsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export service list export downloads default response a status code equal to that given
func (o *ExportServiceListExportDownloadsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export service list export downloads default response
func (o *ExportServiceListExportDownloadsDefault) Code() int {
	return o._statusCode
}

func (o *ExportServiceListExportDownloadsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1beta1/clusters/{clusterId}/exports/{exportId}/downloads][%d] ExportService_ListExportDownloads default %s", o._statusCode, payload)
}

func (o *ExportServiceListExportDownloadsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1beta1/clusters/{clusterId}/exports/{exportId}/downloads][%d] ExportService_ListExportDownloads default %s", o._statusCode, payload)
}

func (o *ExportServiceListExportDownloadsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ExportServiceListExportDownloadsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
