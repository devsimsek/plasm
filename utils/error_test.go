package utils_test

import (
	"testing"

	"github.com/devsimsek/plasm/utils"
)

var err error

func TestError(t *testing.T) {
	err = nil
	utils.Error(err)
}

func TestErrorInfo(t *testing.T) {
	err = nil
	message := "%v"
	args := []interface{}{"Some msg"}
	utils.ErrorInfo(err, message, args...)
}
