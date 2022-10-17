package core

import "errors"

func ChainError(message string, err error) error {
	return errors.New(message + "\n  " + err.Error())
}
