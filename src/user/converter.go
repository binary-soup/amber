package user

import (
	"github.com/binary-soup/amber/core"
	"github.com/binary-soup/amber/core/nodes"
	"golang.org/x/crypto/bcrypt"
)

type converterNode struct{}

func (converterNode) Run(_ nodes.WorkflowContext, input any) (any, error) {
	userInput := input.(*CreateUserInput)

	hash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, core.ChainError("error hashing password", err)
	}

	return NewUser(userInput.Username, hash), nil
}
