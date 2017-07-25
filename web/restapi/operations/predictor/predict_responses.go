// Code generated by go-swagger; DO NOT EDIT.

package predictor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rai-project/dlframework/web/models"
)

// PredictOKCode is the HTTP code returned for type PredictOK
const PredictOKCode int = 200

/*PredictOK predict o k

swagger:response predictOK
*/
type PredictOK struct {

	/*
	  In: Body
	*/
	Payload *models.DlframeworkPredictResponse `json:"body,omitempty"`
}

// NewPredictOK creates PredictOK with default headers values
func NewPredictOK() *PredictOK {
	return &PredictOK{}
}

// WithPayload adds the payload to the predict o k response
func (o *PredictOK) WithPayload(payload *models.DlframeworkPredictResponse) *PredictOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the predict o k response
func (o *PredictOK) SetPayload(payload *models.DlframeworkPredictResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PredictOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
