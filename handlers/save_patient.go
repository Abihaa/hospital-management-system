package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abihaa/hospital-management-system"
	"github.com/abihaa/hospital-management-system/gen/restapi/operations"
)

// NewSavePatient handles request for saving patient record.
func NewSavePatient(rt *runtime.Runtime) operations.SavePatientHandler {
	return &savePatient{rt: rt}
}

type savePatient struct {
	rt *runtime.Runtime
}

// Handle the savePatient request.
func (d *savePatient) Handle(params operations.SavePatientParams) middleware.Responder {
	id, err := d.rt.Service().SavePatient(ToDomainPatient(params.Patient))
	if err != nil {
		log().Errorf("failed to save: params: %+v error: %+v ", params, err)
		return operations.NewSavePatientConflict()
	}

	params.Patient.ID = id
	return operations.NewSavePatientCreated().WithPayload(params.Patient)
}
