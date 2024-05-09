package utils

import "fmt"

func NotSavedError(model, name string, err error) error {
	return fmt.Errorf("%s %s not saved: %v", model, name, err)
}
