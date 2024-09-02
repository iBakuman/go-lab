package reflect_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	var a any
	va := reflect.ValueOf(a)
	require.False(t, va.IsValid())
	vb := reflect.Value{}
	require.False(t, vb.IsValid())

	var c any
	var d int
	c = d
	vc := reflect.ValueOf(c)
	require.True(t, vc.IsValid())

	type Iface interface {
		Get() int
	}
	var e Iface
	var f any
	f = e
	vf := reflect.ValueOf(f)
	require.True(t, vf.IsValid())
}

// Invalid Value Initialization: A [reflect.Value] that is created using reflect.Value{} or any uninitialized value is considered invalid. For example:
func TestZeroValue(t *testing.T) {
	var v reflect.Value
	require.False(t, v.IsValid())
}

// Nil Pointers, Interfaces, and Slices: If a [reflect.Value] represents a nil pointer, nil interface, or nil slice, it is considered invalid:
func TestNilValue(t *testing.T) {
	var a *int
	va := reflect.ValueOf(a)
	require.True(t, va.IsValid())
	require.False(t, va.Elem().IsValid())

	var b interface{}
	vb := reflect.ValueOf(b)
	require.False(t, vb.IsValid())

	var c []int
	vc := reflect.ValueOf(c)
	require.True(t, vc.IsValid())
}

// Out-of-bound Access: If you try to access an element or field that doesn't exist, it results in an invalid reflect.Value:
func TestOutOfBounds(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		type MyStruct struct{}
		v := reflect.ValueOf(MyStruct{})
		v = v.FieldByName("NonExistentField")
		fmt.Println(v.IsValid()) // false, because the field does not exist
	})
	t.Run("slice", func(t *testing.T) {
		a := []int{1, 2, 3}
		va := reflect.ValueOf(a)
		require.True(t, va.IsValid())
		require.Panics(t, func() {
			va.Index(3)
		})
	})

}

// Map Lookup: If a map lookup using reflection doesn't find the key, it returns an invalid reflect.Value:
func TestMapLookup(t *testing.T) {
	m := map[string]int{"one": 1}
	v := reflect.ValueOf(m)
	v = v.MapIndex(reflect.ValueOf("two")) // Key "two" does not exist
	require.False(t, v.IsValid())
}

// Summary
// 1. IsValid() == false:
// This indicates that the [reflect.Value] is invalid, usually meaning it is the zero value of reflect.Value.
// This can occur when dealing with nil values, non-existent fields, out-of-bounds indices, or map lookups that fail to find a key.
// 2. IsValid() == true:
// The reflect.Value is valid and represents a real value, whether it is a pointer, struct, array, map, etc.
// So, IsValid() returning false is specifically tied to cases where the [reflect.Value] does not represent a valid, meaningful Go value.