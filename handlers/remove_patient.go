package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abihaa/hospital-management-system"
	domainErr "github.com/abihaa/hospital-management-system/errors"
	"github.com/abihaa/hospital-management-system/gen/restapi/operations"
)

// NewRemovePatient handles remove patient request.
func NewRemovePatient(rt *runtime.Runtime) operations.RemovePatientHandler {
	return &removePatient{rt: rt}
}

type removePatient struct {
	rt *runtime.Runtime
}

func (r removePatient) Handle(params operations.RemovePatientParams) middleware.Responder {
	if err := r.rt.Service().RemovePatient(params.ID); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			log().Errorf("failed to remove: params: %+v error:%+v ", params, err)
			return operations.NewGetPatientNotFound()
		default:
			log().Errorf("failed to remove: params: %+v error:%+v ", params, err)
			return operations.NewRemovePatientInternalServerError()
		}
	}
	return operations.NewRemovePatientNoContent()
}
