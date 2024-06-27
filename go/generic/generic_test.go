package generic

import "testing"

type Person[T any] struct {
	Name string
}

func (p *Person[T]) Echo(t T) T {
	return t
}

func TestStructGeneric(t *testing.T) {

}
