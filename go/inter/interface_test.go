package inter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Person struct {
	Name string
	Age  int
}

func TestValue(t *testing.T) {
	a := Person{
		Name: "A",
		Age:  1,
	}
	b := a
	a.Name = "B"
	require.Equal(t, "B", a.Name)
	require.Equal(t, "A", b.Name)
}

func TestInterfaceCopyValue(t *testing.T) {
	var a any
	pa := Person{
		Name: "A",
	}
	a = pa
	// Not allowed
	// a.(Person).Name = "John"
	a = &Person{}
	// Allowed
	a.(*Person).Name = "John"
	require.Equal(t, "John", a.(*Person).Name)
}

func TestTypeAssert(t *testing.T) {
	var a any
	require.Equal(t, nil, a)
	require.Panics(t, func() {
		_ = a.(*Person)
	})
}
