package multierr_test

import (
	"errors"
	"testing"

	"github.com/hashicorp/go-multierror"
)

func TestAppend(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	errs := multierror.Append(err1, err2)
	t.Logf("%+v", errs)
	t.Logf("%v", errs.Errors)
}
