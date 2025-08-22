//go:build test

package testkit

import (
	"context"
	"fmt"
	"github.com/ibakuman/go-lab/go/pkgtmp"
)

var _ pkgtmp.Handler = &hImpl{}

type hImpl struct{}

func (h *hImpl) Handle(ctx context.Context, data []byte) error {
	fmt.Println(string(data))
	return nil
}

func NewHandler() pkgtmp.Handler {
	return &hImpl{}
}