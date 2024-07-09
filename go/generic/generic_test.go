package generic

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints"
)

type Person[T any] struct {
	Name string
}

func (p *Person[T]) Echo(t T) T {
	return t
}

func TestStructGeneric(t *testing.T) {

}

// func Scale[E constraints.Integer](s []E, c E) []E {
// 	r := make([]E, len(s))
// 	for i, v := range s {
// 		r[i] = v * c
// 	}
// 	return r
// }
//
// type Point []int32
//
// func (p Point) String() string {
// 	// Details not important.
// 	return "point"
// }
//
// // ScaleAndPrint doubles a Point and prints it.
// func ScaleAndPrint(p Point) {
// 	r := Scale(p, 2)
// 	fmt.Println(r.String()) // DOES NOT COMPILE
// }

func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

type Point []int32

func (p Point) String() string {
	// Details not important.
	return "point"
}

// ScaleAndPrint doubles a Point and prints it.
func ScaleAndPrint(p Point) {
	r := Scale(p, 2)
	fmt.Println(r.String()) // DOES NOT COMPILE
}
func TestConstraintTypeInference(t *testing.T) {

}
