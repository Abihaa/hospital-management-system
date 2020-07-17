// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListPatientHandlerFunc turns a function with the right signature into a list patient handler
type ListPatientHandlerFunc func(ListPatientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListPatientHandlerFunc) Handle(params ListPatientParams) middleware.Responder {
	return fn(params)
}

// ListPatientHandler interface for that can handle valid list patient params
type ListPatientHandler interface {
	Handle(ListPatientParams) middleware.Responder
}

// NewListPatient creates a new http.Handler for the list patient operation
func NewListPatient(ctx *middleware.Context, handler ListPatientHandler) *ListPatient {
	return &ListPatient{Context: ctx, Handler: handler}
}

/*ListPatient swagger:route GET /patient listPatient

return patients records based on filters

*/
type ListPatient struct {
	Context *middleware.Context
	Handler ListPatientHandler
}

func (o *ListPatient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListPatientParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
