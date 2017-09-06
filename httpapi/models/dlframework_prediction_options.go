// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkPredictionOptions dlframework prediction options
// swagger:model dlframeworkPredictionOptions

type DlframeworkPredictionOptions struct {

	// feature limit
	FeatureLimit int32 `json:"feature_limit,omitempty"`

	// request id
	RequestID string `json:"request_id,omitempty"`
}

/* polymorph dlframeworkPredictionOptions feature_limit false */

/* polymorph dlframeworkPredictionOptions request_id false */

// Validate validates this dlframework prediction options
func (m *DlframeworkPredictionOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkPredictionOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkPredictionOptions) UnmarshalBinary(b []byte) error {
	var res DlframeworkPredictionOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
