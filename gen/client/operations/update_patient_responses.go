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

// UpdatePatientReader is a Reader for the UpdatePatient structure.
type UpdatePatientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePatientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePatientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdatePatientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePatientOK creates a UpdatePatientOK with default headers values
func NewUpdatePatientOK() *UpdatePatientOK {
	return &UpdatePatientOK{}
}

/*UpdatePatientOK handles this case with default header values.

patient updated
*/
type UpdatePatientOK struct {
	Payload *models.Patient
}

func (o *UpdatePatientOK) Error() string {
	return fmt.Sprintf("[PUT /patient/{ID}][%d] updatePatientOK  %+v", 200, o.Payload)
}

func (o *UpdatePatientOK) GetPayload() *models.Patient {
	return o.Payload
}

func (o *UpdatePatientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Patient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePatientInternalServerError creates a UpdatePatientInternalServerError with default headers values
func NewUpdatePatientInternalServerError() *UpdatePatientInternalServerError {
	return &UpdatePatientInternalServerError{}
}

/*UpdatePatientInternalServerError handles this case with default header values.

internal server error
*/
type UpdatePatientInternalServerError struct {
}

func (o *UpdatePatientInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /patient/{ID}][%d] updatePatientInternalServerError ", 500)
}

func (o *UpdatePatientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
