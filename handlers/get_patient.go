package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abihaa/hospital-management-system"
	domainErr "github.com/abihaa/hospital-management-system/errors"
	"github.com/abihaa/hospital-management-system/gen/restapi/operations"
)

// NewGetPatient handles request for retrieving data.
func NewGetPatient(rt *runtime.Runtime) operations.GetPatientHandler {
	return &getPatient{rt: rt}
}

type getPatient struct {
	rt *runtime.Runtime
}

func (g *getPatient) Handle(params operations.GetPatientParams) middleware.Responder {
	patient, err := g.rt.Service().GetPatientByID(params.ID)
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			log().Errorf("failed to get patient: params: %+v error:%+v ", params, err)
			return operations.NewGetPatientNotFound()
		default:
			log().Errorf("failed to get patient: params: %+v error:%+v ", params, err)
			return operations.NewGetPatientInternalServerError()
		}
	}
	return operations.NewGetPatientOK().WithPayload(ToGenPatient(patient))
}
