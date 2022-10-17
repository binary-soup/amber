package web

import (
	"net/http"

	"github.com/binary-soup/amber/core/nodes"
)

type DataResponseNode struct{}

func (DataResponseNode) Run(ctx nodes.WorkflowContext, input any) (any, error) {
	sendJSONResponse(ctx.ResponseWriter, http.StatusOK, NewSuccessDataResponse(input))
	return nil, nil
}
