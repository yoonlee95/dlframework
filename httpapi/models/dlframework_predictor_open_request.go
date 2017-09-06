// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkPredictorOpenRequest dlframework predictor open request
// swagger:model dlframeworkPredictorOpenRequest

type DlframeworkPredictorOpenRequest struct {

	// framework name
	FrameworkName string `json:"framework_name,omitempty"`

	// framework version
	FrameworkVersion string `json:"framework_version,omitempty"`

	// model name
	ModelName string `json:"model_name,omitempty"`

	// model version
	ModelVersion string `json:"model_version,omitempty"`
}

/* polymorph dlframeworkPredictorOpenRequest framework_name false */

/* polymorph dlframeworkPredictorOpenRequest framework_version false */

/* polymorph dlframeworkPredictorOpenRequest model_name false */

/* polymorph dlframeworkPredictorOpenRequest model_version false */

// Validate validates this dlframework predictor open request
func (m *DlframeworkPredictorOpenRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkPredictorOpenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkPredictorOpenRequest) UnmarshalBinary(b []byte) error {
	var res DlframeworkPredictorOpenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
