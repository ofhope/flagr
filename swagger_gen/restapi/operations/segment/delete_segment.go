// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteSegmentHandlerFunc turns a function with the right signature into a delete segment handler
type DeleteSegmentHandlerFunc func(DeleteSegmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSegmentHandlerFunc) Handle(params DeleteSegmentParams) middleware.Responder {
	return fn(params)
}

// DeleteSegmentHandler interface for that can handle valid delete segment params
type DeleteSegmentHandler interface {
	Handle(DeleteSegmentParams) middleware.Responder
}

// NewDeleteSegment creates a new http.Handler for the delete segment operation
func NewDeleteSegment(ctx *middleware.Context, handler DeleteSegmentHandler) *DeleteSegment {
	return &DeleteSegment{Context: ctx, Handler: handler}
}

/*
	DeleteSegment swagger:route DELETE /flags/{flagID}/segments/{segmentID} segment deleteSegment

DeleteSegment delete segment API
*/
type DeleteSegment struct {
	Context *middleware.Context
	Handler DeleteSegmentHandler
}

func (o *DeleteSegment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteSegmentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
