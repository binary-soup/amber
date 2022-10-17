package user

import (
	"errors"
	"regexp"
	"strings"

	"github.com/binary-soup/amber/core/nodes"
)

type validatorNode struct{}

func (validatorNode) Run(_ nodes.WorkflowContext, input any) (any, error) {
	user := input.(*CreateUserInput)
	var errs []string

	if len(user.Username) < 5 {
		errs = append(errs, "username shorter than min length 5")
	} else if len(user.Username) > 30 {
		errs = append(errs, "username longer than max length 30")
	}

	if len(user.Password) < 8 {
		errs = append(errs, "password shorter than min length 8")
	}

	if regexp.MustCompile(`[0-9]`).FindString(user.Password) == "" {
		errs = append(errs, "password must contain digit")
	}

	if regexp.MustCompile(`[^0-9a-zA-Z]`).FindString(user.Password) == "" {
		errs = append(errs, "password must contain symbol")
	}

	if len(errs) > 0 {
		return user, errors.New(strings.Join(errs, ", "))
	}

	return user, nil
}
