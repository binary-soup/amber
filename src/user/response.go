package user

import "github.com/binary-soup/amber/core/nodes"

type UserResponse struct {
	Username string `json:"username"`
}

type userResponseConverterNode struct{}

func (userResponseConverterNode) Run(_ nodes.WorkflowContext, input any) (any, error) {
	user := input.(*User)

	return UserResponse{
		Username: user.Username,
	}, nil
}
