package user

import (
	"github.com/binary-soup/amber/core/nodes"
	"github.com/binary-soup/amber/core/nodes/web"
)

type CreateUserInput struct {
	Username string
	Password string
}

func CreateWorkflow() nodes.Workflow {
	return nodes.NewWorkflow(
		validatorNode{},
		converterNode{},
		//save to database
	)
}

func CreateEndpoint() nodes.Workflow {
	return nodes.NewWorkflow(
		web.JSONBodyParserNode[CreateUserInput]{},
		CreateWorkflow(),
		userResponseConverterNode{},
		web.DataResponseNode{},
	)
}
