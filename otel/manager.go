package main

import (
	"github.com/ibakuman/go-lab/otel/embedded"
	"github.com/ibakuman/go-lab/otel/manager"
)

type ManagerProvider struct {
	embedded.ManagerProvider
}

func (p *ManagerProvider) Manager() string {
	return "manager"
}

func Foo() {
	var a manager.Provider = &ManagerProvider{}
	s := a.Manager()
	_ = s
}
