// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetModelManifestsHandlerFunc turns a function with the right signature into a get model manifests handler
type GetModelManifestsHandlerFunc func(GetModelManifestsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetModelManifestsHandlerFunc) Handle(params GetModelManifestsParams) middleware.Responder {
	return fn(params)
}

// GetModelManifestsHandler interface for that can handle valid get model manifests params
type GetModelManifestsHandler interface {
	Handle(GetModelManifestsParams) middleware.Responder
}

// NewGetModelManifests creates a new http.Handler for the get model manifests operation
func NewGetModelManifests(ctx *middleware.Context, handler GetModelManifestsHandler) *GetModelManifests {
	return &GetModelManifests{Context: ctx, Handler: handler}
}

/*GetModelManifests swagger:route GET /v1/models Registry getModelManifests

GetModelManifests get model manifests API

*/
type GetModelManifests struct {
	Context *middleware.Context
	Handler GetModelManifestsHandler
}

func (o *GetModelManifests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetModelManifestsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
