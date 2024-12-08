package errors

import (
	"testing"

	"github.com/pkg/errors"
)

type ValueError struct{
	msg string
}

func (v ValueError) Error() string {
	return v.msg
}

type PointerError struct{
	msg string
}

func (p *PointerError) Error() string {
	return p.msg
}

func TestAs(t *testing.T) {
	a1 := ValueError{msg: "value error"}
	b1 := PointerError{msg: "pointer error"}

	var a1Err error = a1
	var a2Err error = &b1
	if err := new(ValueError); errors.As(a1Err, &err) {
		t.Logf("A")
	}

	if err := new(ValueError); errors.As(a2Err, &err) {
		t.Logf("B")
	}

	vpErr := &ValueError{msg: "vp error"}
	pppErr := &PointerError{msg: "pp error"}
	if err := new(ValueError); errors.As(vpErr, &err) {
		t.Logf("C")
	}
	if err := new(PointerError); errors.As(pppErr, &err) {
		t.Logf("D")
	}
}