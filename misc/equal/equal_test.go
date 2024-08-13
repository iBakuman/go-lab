// Quoted from https://medium.com/golangspec/equality-in-golang-ff44da79b7f1
package equal_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// If dynamic types are identical and dynamic values are equal then two interface values are equal:
type A int
type B = A
type C int
type I interface{ m() }

func (a A) m() {}
func (c C) m() {}
func TestEqual2(t *testing.T) {
	var a I = A(1)
	var b I = B(1)
	var c I = C(1)
	require.True(t, a == b)
	require.False(t, a == c)
	require.False(t, b == c)
}

// It’s possible to compare value x of non-interface type X with value i of interface type I.
// There are few limitations though:
// 1. type X implements interface I
// 2. type X is comparable
// If dynamic type of i is X and dynamic value of i is equal to x then values are equal (source code):

type X int

func (x X) m() {}

type Y int

func (y Y) m() {}

type Z int

func TestEqual3(t *testing.T) {
	var i I = X(1)
	require.True(t, i == X(1))
	require.False(t, i == Y(1))
	// fmt.Println(i == Z(1)) // mismatched types I and C
	// fmt.Println(i == 1) // mismatched types I and int
}

// If dynamic types of interface values are identical but not comparable then
// it will generate runtime panic (source code):
type unComparableA []byte

func TestEqual4(t *testing.T) {
	var i interface{} = unComparableA{}
	var j interface{} = unComparableA{}
	require.Panics(t, func() {
		fmt.Println(i == j)
	})
}

// If types are different but still not comparable then interface values
// aren’t equal, but no panic is generated:
type unComparableB []byte
type unComparableC []byte

func TestEqual5(t *testing.T) {
	var i interface{} = unComparableB{}
	var j interface{} = unComparableC{}
	require.False(t, i == j)
}

// ## Structs
// While comparing structs, corresponding non-blank fields are checked for equality — both exported and non-exported (source code):

type StructA struct {
	b  float64
	f1 int
	F2 string
}
type StructB struct {
	_  float64
	f1 int
	F2 string
}

func TestEqual6(t *testing.T) {
	fmt.Println(StructA{1.1, 2, "x"} == StructA{0.1, 2, "x"}) // true
	// fmt.Println(A{} == B{}) // mismatched types A and B
}

// It’s worth to introduce now main rule applicable not only for structs but all types:
// x == y is allowed only when either x is assignable to y or y is assignable to x.
// This is why A{} == B{} above generates compile-time error.

// ## Arrays
// This is similar to struct explained earlier. Corresponding elements needs to be equal
// for the whole arrays to be equal:

type T struct {
	name string
	age  int
	_    float64
}

func TestEqual7(t *testing.T) {
	x := [...]float64{1.1, 2, 3.14}
	require.True(t, x == [...]float64{1.1, 2, 3.14})
	// y := [1]T{{"foo", 1, 0}} // Cannot assign a value to a blank field
	y := [1]T{{name: "foo", age: 1}}
	require.True(t, y == [1]T{{name: "foo", age: 1}})
}

// ## Comparing something non-comparable
// In this bucket we’ve three types: functions, maps and slices.
// We can’t do much about the functions. There is not way to compare them in Go (source code):

func TestEqual8(t *testing.T) {
	f := func(int) int { return 1 }
	g := func(int) int { return 2 }
	_, _ = f, g
	// f == g // It generates compile-time error:
	// invalid operation: f == g (func can only be compared to nil).
	// It also gives a hint that we can compare functions to nil.
	// The same is true for maps and slices
	f1 := func(int) int { return 1 }
	m1 := make(map[int]int)
	s1 := make([]byte, 10)
	require.False(t, f1 == nil)
	require.False(t, m1 == nil)
	require.False(t, s1 == nil)
	var f2 func()
	var m2 map[int]int
	var s2 []byte
	require.True(t, f2 == nil)
	require.True(t, m2 == nil)
	require.True(t, s2 == nil)
}

// Are there any options for maps or slices though? Luckily there are and we’ll explore them right now…

// ## []byte
// Package bytes offers utilities to deal with byte slices, and it provides functions to
// check if slices are equal and even equal under Unicode case-folding (source code):

func TestEqual9(t *testing.T) {
	s1 := []byte{'f', 'o', 'o'}
	s2 := []byte{'f', 'o', 'o'}
	require.True(t, bytes.Equal(s1, s2))
	s2 = []byte{'b', 'a', 'r'}
	require.False(t, bytes.Equal(s1, s2))
	s2 = []byte{'f', 'O', 'O'}
	require.True(t, bytes.EqualFold(s1, s2))
	s1 = []byte("źdźbło")
	s2 = []byte("źdŹbŁO")
	require.True(t, bytes.EqualFold(s1, s2))
	s1 = []byte{}
	s2 = nil
	// Be aware the following comparison is ok
	require.True(t, bytes.Equal(s1, s2))
}

// What about maps or slices where elements of underlying arrays are not bytes?
// We’ve two options: reflect.DeepEqual , cmp package or writing ad-hoc comparison code using
// e.g. for statement. Let’s see first two approaches in action.
