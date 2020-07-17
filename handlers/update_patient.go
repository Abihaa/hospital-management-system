package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abihaa/hospital-management-system"
	"github.com/abihaa/hospital-management-system/gen/restapi/operations"
)

// NewUpdatePatient handles update request.
func NewUpdatePatient(rt *runtime.Runtime) operations.UpdatePatientHandler {
	return &updatePatient{rt: rt}
}

type updatePatient struct {
	rt *runtime.Runtime
}

func (u *updatePatient) Handle(params operations.UpdatePatientParams) middleware.Responder {
	_, err := u.rt.Service().GetPatientByID(params.ID)
	if err != nil {
		log().Errorf("failed to update: params: %+v error:%+v ", params, err)
		return operations.NewGetPatientNotFound()
	}
	params.Patient.ID = params.ID
	if err := u.rt.Service().UpdatePatient(ToDomainPatient(params.Patient)); err != nil {
		log().Errorf("failed to update: params: %+v error:%+v ", params, err)
		return operations.NewUpdatePatientInternalServerError()
	}

	return operations.NewUpdatePatientOK()
}
