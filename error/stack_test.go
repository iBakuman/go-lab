package error

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func A() error {
	return B()
}

func B() error {
	return C()
}

func C() error {
	return errors.New("error")
}

func TestErrorStack(t *testing.T) {
	err := A()
	fmt.Printf("%+v\n", err)
}
