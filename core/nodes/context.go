package nodes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type WorkflowContext struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         httprouter.Params
}

func NewWorkflowContext(w http.ResponseWriter, req *http.Request, params httprouter.Params) WorkflowContext {
	return WorkflowContext{
		ResponseWriter: w,
		Request:        req,
		Params:         params,
	}
}
