// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SavePatientHandlerFunc turns a function with the right signature into a save patient handler
type SavePatientHandlerFunc func(SavePatientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SavePatientHandlerFunc) Handle(params SavePatientParams) middleware.Responder {
	return fn(params)
}

// SavePatientHandler interface for that can handle valid save patient params
type SavePatientHandler interface {
	Handle(SavePatientParams) middleware.Responder
}

// NewSavePatient creates a new http.Handler for the save patient operation
func NewSavePatient(ctx *middleware.Context, handler SavePatientHandler) *SavePatient {
	return &SavePatient{Context: ctx, Handler: handler}
}

/*SavePatient swagger:route POST /patient savePatient

SavePatient save patient API

*/
type SavePatient struct {
	Context *middleware.Context
	Handler SavePatientHandler
}

func (o *SavePatient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSavePatientParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
