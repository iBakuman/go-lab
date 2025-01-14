package manager

import (
	"github.com/ibakuman/go-lab/otel/embedded"
)

type Provider interface{
	embedded.ManagerProvider
	Manager() string
}