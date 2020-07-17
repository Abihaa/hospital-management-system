// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/abihaa/hospital-management-system/gen/models"
)

// ListPatientReader is a Reader for the ListPatient structure.
type ListPatientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPatientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPatientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListPatientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPatientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPatientOK creates a ListPatientOK with default headers values
func NewListPatientOK() *ListPatientOK {
	return &ListPatientOK{}
}

/*ListPatientOK handles this case with default header values.

patient found
*/
type ListPatientOK struct {
	Payload []*models.Patient
}

func (o *ListPatientOK) Error() string {
	return fmt.Sprintf("[GET /patient][%d] listPatientOK  %+v", 200, o.Payload)
}

func (o *ListPatientOK) GetPayload() []*models.Patient {
	return o.Payload
}

func (o *ListPatientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPatientNotFound creates a ListPatientNotFound with default headers values
func NewListPatientNotFound() *ListPatientNotFound {
	return &ListPatientNotFound{}
}

/*ListPatientNotFound handles this case with default header values.

patient not found
*/
type ListPatientNotFound struct {
}

func (o *ListPatientNotFound) Error() string {
	return fmt.Sprintf("[GET /patient][%d] listPatientNotFound ", 404)
}

func (o *ListPatientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListPatientInternalServerError creates a ListPatientInternalServerError with default headers values
func NewListPatientInternalServerError() *ListPatientInternalServerError {
	return &ListPatientInternalServerError{}
}

/*ListPatientInternalServerError handles this case with default header values.

internal server error
*/
type ListPatientInternalServerError struct {
}

func (o *ListPatientInternalServerError) Error() string {
	return fmt.Sprintf("[GET /patient][%d] listPatientInternalServerError ", 500)
}

func (o *ListPatientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
