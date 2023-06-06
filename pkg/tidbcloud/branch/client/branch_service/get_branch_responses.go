// Code generated by go-swagger; DO NOT EDIT.

package branch_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"tidbcloud-cli/pkg/tidbcloud/branch/models"
)

// GetBranchReader is a Reader for the GetBranch structure.
type GetBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBranchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBranchOK creates a GetBranchOK with default headers values
func NewGetBranchOK() *GetBranchOK {
	return &GetBranchOK{}
}

/*
GetBranchOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetBranchOK struct {
	Payload *models.OpenapiGetBranchResp
}

// IsSuccess returns true when this get branch o k response has a 2xx status code
func (o *GetBranchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get branch o k response has a 3xx status code
func (o *GetBranchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get branch o k response has a 4xx status code
func (o *GetBranchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get branch o k response has a 5xx status code
func (o *GetBranchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get branch o k response a status code equal to that given
func (o *GetBranchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get branch o k response
func (o *GetBranchOK) Code() int {
	return 200
}

func (o *GetBranchOK) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/clusters/{cluster_id}/branches/{branch_id}][%d] getBranchOK  %+v", 200, o.Payload)
}

func (o *GetBranchOK) String() string {
	return fmt.Sprintf("[GET /api/v1beta/clusters/{cluster_id}/branches/{branch_id}][%d] getBranchOK  %+v", 200, o.Payload)
}

func (o *GetBranchOK) GetPayload() *models.OpenapiGetBranchResp {
	return o.Payload
}

func (o *GetBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenapiGetBranchResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBranchDefault creates a GetBranchDefault with default headers values
func NewGetBranchDefault(code int) *GetBranchDefault {
	return &GetBranchDefault{
		_statusCode: code,
	}
}

/*
GetBranchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetBranchDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this get branch default response has a 2xx status code
func (o *GetBranchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get branch default response has a 3xx status code
func (o *GetBranchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get branch default response has a 4xx status code
func (o *GetBranchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get branch default response has a 5xx status code
func (o *GetBranchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get branch default response a status code equal to that given
func (o *GetBranchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get branch default response
func (o *GetBranchDefault) Code() int {
	return o._statusCode
}

func (o *GetBranchDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1beta/clusters/{cluster_id}/branches/{branch_id}][%d] GetBranch default  %+v", o._statusCode, o.Payload)
}

func (o *GetBranchDefault) String() string {
	return fmt.Sprintf("[GET /api/v1beta/clusters/{cluster_id}/branches/{branch_id}][%d] GetBranch default  %+v", o._statusCode, o.Payload)
}

func (o *GetBranchDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetBranchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}