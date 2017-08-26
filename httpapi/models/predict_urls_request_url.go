// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PredictUrlsRequestURL predict urls request URL
// swagger:model PredictURLsRequestURL
type PredictUrlsRequestURL struct {

	// An id used to identify the output feature: maps to input_id for output
	ID string `json:"id,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this predict urls request URL
func (m *PredictUrlsRequestURL) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PredictUrlsRequestURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PredictUrlsRequestURL) UnmarshalBinary(b []byte) error {
	var res PredictUrlsRequestURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}