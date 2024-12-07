package reflect

import (
	"reflect"
	"runtime"
	"testing"
)

type PersonFunc struct {
	Name string
}

func (p *PersonFunc) GetName() string {
	return p.Name
}

type TestInterfaceA interface {
	GetName() string
}

func TestFuncPtr(t *testing.T) {
	p := PersonFunc{Name: "Alice"}
	v := reflect.ValueOf(p.GetName)
	t.Logf("v.Type: %v", v.Type())
	t.Logf("v.Kind: %v", v.Kind())
	funcName := runtime.FuncForPC(v.Pointer())
	t.Logf("funcName: %v", funcName.Name())
}
