package unsafe

import (
	"testing"
	"unsafe"
)

func modifyInteger(num *int) {
	*num = 100
}

func modifyString(str *string) {
	*str = "world"
}

func TestUnsafe_Pointer(t *testing.T) {
	var i int = 32
	var f *float64
	f = (*float64)(unsafe.Pointer(&i))
	t.Log(*f)
	t.Log("before modify i:", i)
	modifyInteger(&i)
	t.Log("after modify i:", i)
	var s string = "hello"
	t.Log("before modify s:", s)
	modifyString(&s)
	t.Log("after modify s:", s)
}
