// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// ListPatientURL generates an URL for the list patient operation
type ListPatientURL struct {
	Age        *int64
	Conditions *string
	Gender     *string
	Limit      *int64
	Name       *string
	Offset     *int64
	Phone      *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListPatientURL) WithBasePath(bp string) *ListPatientURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListPatientURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ListPatientURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/patient"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var ageQ string
	if o.Age != nil {
		ageQ = swag.FormatInt64(*o.Age)
	}
	if ageQ != "" {
		qs.Set("age", ageQ)
	}

	var conditionsQ string
	if o.Conditions != nil {
		conditionsQ = *o.Conditions
	}
	if conditionsQ != "" {
		qs.Set("conditions", conditionsQ)
	}

	var genderQ string
	if o.Gender != nil {
		genderQ = *o.Gender
	}
	if genderQ != "" {
		qs.Set("gender", genderQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var nameQ string
	if o.Name != nil {
		nameQ = *o.Name
	}
	if nameQ != "" {
		qs.Set("name", nameQ)
	}

	var offsetQ string
	if o.Offset != nil {
		offsetQ = swag.FormatInt64(*o.Offset)
	}
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	var phoneQ string
	if o.Phone != nil {
		phoneQ = *o.Phone
	}
	if phoneQ != "" {
		qs.Set("phone", phoneQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ListPatientURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ListPatientURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ListPatientURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ListPatientURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ListPatientURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ListPatientURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
