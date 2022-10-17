package web

import (
	"net/http"

	"github.com/binary-soup/amber/core/nodes"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	Workflow nodes.Workflow
}

func NewHandler(w nodes.Workflow) Handler {
	return Handler{
		Workflow: w,
	}
}

func (h Handler) Serve(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := nodes.NewWorkflowContext(w, req, params)

	_, err := h.Workflow.Run(ctx, nil)
	if err != nil {
		sendErrorResponse(w, err)
	}
}
