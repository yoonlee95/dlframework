// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rai-project/dlframework/httpapi/models"
)

// FrameworkManifestsOKCode is the HTTP code returned for type FrameworkManifestsOK
const FrameworkManifestsOKCode int = 200

/*FrameworkManifestsOK framework manifests o k

swagger:response frameworkManifestsOK
*/
type FrameworkManifestsOK struct {

	/*
	  In: Body
	*/
	Payload *models.DlframeworkFrameworkManifestsResponse `json:"body,omitempty"`
}

// NewFrameworkManifestsOK creates FrameworkManifestsOK with default headers values
func NewFrameworkManifestsOK() *FrameworkManifestsOK {
	return &FrameworkManifestsOK{}
}

// WithPayload adds the payload to the framework manifests o k response
func (o *FrameworkManifestsOK) WithPayload(payload *models.DlframeworkFrameworkManifestsResponse) *FrameworkManifestsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the framework manifests o k response
func (o *FrameworkManifestsOK) SetPayload(payload *models.DlframeworkFrameworkManifestsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FrameworkManifestsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
