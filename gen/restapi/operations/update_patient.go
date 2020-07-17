// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdatePatientHandlerFunc turns a function with the right signature into a update patient handler
type UpdatePatientHandlerFunc func(UpdatePatientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdatePatientHandlerFunc) Handle(params UpdatePatientParams) middleware.Responder {
	return fn(params)
}

// UpdatePatientHandler interface for that can handle valid update patient params
type UpdatePatientHandler interface {
	Handle(UpdatePatientParams) middleware.Responder
}

// NewUpdatePatient creates a new http.Handler for the update patient operation
func NewUpdatePatient(ctx *middleware.Context, handler UpdatePatientHandler) *UpdatePatient {
	return &UpdatePatient{Context: ctx, Handler: handler}
}

/*UpdatePatient swagger:route PUT /patient/{ID} updatePatient

UpdatePatient update patient API

*/
type UpdatePatient struct {
	Context *middleware.Context
	Handler UpdatePatientHandler
}

func (o *UpdatePatient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdatePatientParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
