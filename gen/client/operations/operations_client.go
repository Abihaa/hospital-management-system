// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetPatient(params *GetPatientParams) (*GetPatientOK, error)

	ListPatient(params *ListPatientParams) (*ListPatientOK, error)

	RemovePatient(params *RemovePatientParams) (*RemovePatientNoContent, error)

	SavePatient(params *SavePatientParams) (*SavePatientCreated, error)

	UpdatePatient(params *UpdatePatientParams) (*UpdatePatientOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetPatient return patient based on UUID
*/
func (a *Client) GetPatient(params *GetPatientParams) (*GetPatientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPatientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPatient",
		Method:             "GET",
		PathPattern:        "/patient/{ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPatientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPatientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPatient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPatient return patients records based on filters
*/
func (a *Client) ListPatient(params *ListPatientParams) (*ListPatientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPatientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPatient",
		Method:             "GET",
		PathPattern:        "/patient",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPatientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPatientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPatient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemovePatient remove patient API
*/
func (a *Client) RemovePatient(params *RemovePatientParams) (*RemovePatientNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemovePatientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removePatient",
		Method:             "DELETE",
		PathPattern:        "/patient/{ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemovePatientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemovePatientNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removePatient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SavePatient save patient API
*/
func (a *Client) SavePatient(params *SavePatientParams) (*SavePatientCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSavePatientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "savePatient",
		Method:             "POST",
		PathPattern:        "/patient",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SavePatientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SavePatientCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for savePatient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePatient update patient API
*/
func (a *Client) UpdatePatient(params *UpdatePatientParams) (*UpdatePatientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePatientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePatient",
		Method:             "PUT",
		PathPattern:        "/patient/{ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePatientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePatientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePatient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
