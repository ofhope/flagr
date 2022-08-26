// Code generated by go-swagger; DO NOT EDIT.

package evaluation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostEvaluationBatchHandlerFunc turns a function with the right signature into a post evaluation batch handler
type PostEvaluationBatchHandlerFunc func(PostEvaluationBatchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostEvaluationBatchHandlerFunc) Handle(params PostEvaluationBatchParams) middleware.Responder {
	return fn(params)
}

// PostEvaluationBatchHandler interface for that can handle valid post evaluation batch params
type PostEvaluationBatchHandler interface {
	Handle(PostEvaluationBatchParams) middleware.Responder
}

// NewPostEvaluationBatch creates a new http.Handler for the post evaluation batch operation
func NewPostEvaluationBatch(ctx *middleware.Context, handler PostEvaluationBatchHandler) *PostEvaluationBatch {
	return &PostEvaluationBatch{Context: ctx, Handler: handler}
}

/*
	PostEvaluationBatch swagger:route POST /evaluation/batch evaluation postEvaluationBatch

PostEvaluationBatch post evaluation batch API
*/
type PostEvaluationBatch struct {
	Context *middleware.Context
	Handler PostEvaluationBatchHandler
}

func (o *PostEvaluationBatch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostEvaluationBatchParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
