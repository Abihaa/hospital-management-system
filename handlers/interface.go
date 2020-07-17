package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/abihaa/hospital-management-system"
	"github.com/abihaa/hospital-management-system/gen/restapi/operations"
)

// Handler replaces swagger handler.
type Handler *operations.HospitalManagementSystemAPI

// NewHandler overrides swagger api handlers.
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewHospitalManagementSystemAPI(spec)

	// patient handlers
	handler.SavePatientHandler = NewSavePatient(rt)
	handler.GetPatientHandler = NewGetPatient(rt)
	handler.UpdatePatientHandler = NewUpdatePatient(rt)
	handler.RemovePatientHandler = NewRemovePatient(rt)
	handler.ListPatientHandler = NewListPatient(rt)

	return handler
}
