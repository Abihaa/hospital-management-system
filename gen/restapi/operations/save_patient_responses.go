// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/abihaa/hospital-management-system/gen/models"
)

// SavePatientCreatedCode is the HTTP code returned for type SavePatientCreated
const SavePatientCreatedCode int = 201

/*SavePatientCreated patient added

swagger:response savePatientCreated
*/
type SavePatientCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Patient `json:"body,omitempty"`
}

// NewSavePatientCreated creates SavePatientCreated with default headers values
func NewSavePatientCreated() *SavePatientCreated {

	return &SavePatientCreated{}
}

// WithPayload adds the payload to the save patient created response
func (o *SavePatientCreated) WithPayload(payload *models.Patient) *SavePatientCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save patient created response
func (o *SavePatientCreated) SetPayload(payload *models.Patient) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SavePatientCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SavePatientConflictCode is the HTTP code returned for type SavePatientConflict
const SavePatientConflictCode int = 409

/*SavePatientConflict patient already exists

swagger:response savePatientConflict
*/
type SavePatientConflict struct {
}

// NewSavePatientConflict creates SavePatientConflict with default headers values
func NewSavePatientConflict() *SavePatientConflict {

	return &SavePatientConflict{}
}

// WriteResponse to the client
func (o *SavePatientConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// SavePatientInternalServerErrorCode is the HTTP code returned for type SavePatientInternalServerError
const SavePatientInternalServerErrorCode int = 500

/*SavePatientInternalServerError internal server error

swagger:response savePatientInternalServerError
*/
type SavePatientInternalServerError struct {
}

// NewSavePatientInternalServerError creates SavePatientInternalServerError with default headers values
func NewSavePatientInternalServerError() *SavePatientInternalServerError {

	return &SavePatientInternalServerError{}
}

// WriteResponse to the client
func (o *SavePatientInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
