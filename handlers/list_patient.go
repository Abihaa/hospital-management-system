package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/abihaa/hospital-management-system"
	domainErr "github.com/abihaa/hospital-management-system/errors"
	genModels "github.com/abihaa/hospital-management-system/gen/models"
	"github.com/abihaa/hospital-management-system/gen/restapi/operations"
)

type listPatient struct {
	rt *runtime.Runtime
}

// NewListPatient handles list patient request.
func NewListPatient(rt *runtime.Runtime) operations.ListPatientHandler {
	return &listPatient{rt: rt}
}

func (l listPatient) Handle(params operations.ListPatientParams) middleware.Responder {
	filter := make(map[string]interface{})
	if params.Name != nil {
		filter["name"] = params.Name
	}
	if params.Age != nil {
		filter["age"] = params.Age
	}
	if params.Gender != nil {
		filter["gender"] = params.Gender
	}
	if params.Phone != nil {
		filter["phone"] = params.Phone
	}
	if params.Conditions != nil {
		filter["conditions"] = params.Conditions
	}
	patient, err := l.rt.Service().ListPatient(filter, swag.Int64Value(params.Offset), swag.Int64Value(params.Limit))
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			log().Errorf("failed to list patient records: params: %+v error:%+v ", params, err)
			return operations.NewGetPatientNotFound()
		default:
			log().Errorf("failed to list patient records: params: %+v error:%+v ", params, err)
			return operations.NewListPatientInternalServerError()
		}
	}

	patientList := make([]*genModels.Patient, 0)
	for _, pat := range patient {
		patientList = append(patientList, ToGenPatient(pat))
	}
	return operations.NewListPatientOK().WithPayload(patientList)
}
