package web

import (
	"encoding/json"
	"errors"

	"github.com/binary-soup/amber/core/nodes"
)

type JSONBodyParserNode[Model any] struct{}

func (JSONBodyParserNode[Model]) Run(ctx nodes.WorkflowContext, input any) (any, error) {
	model := new(Model)

	err := json.NewDecoder(ctx.Request.Body).Decode(model)
	if err != nil {
		return nil, errors.New("invalid json body")
	}

	return model, nil
}
