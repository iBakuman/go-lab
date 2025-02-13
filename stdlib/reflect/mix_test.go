package reflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type InterfaceA interface {
	Hello() string
}

// InterfaceB is an interface that has the same method as InterfaceA
type InterfaceB interface {
	Hello() string
}

type ImplA struct {
}

func (m ImplA) Hello() string {
	return "A"
}

type ImplB struct {
}

func (m ImplB) Hello() string {
	return "B"
}

func TestMix(t *testing.T) {
	ptrIfaceAType := reflect.TypeOf((*InterfaceA)(nil))
	ptrIfaceBType := reflect.TypeOf((*InterfaceB)(nil))
	// It makes sense that ptrIfaceAType and ptrIfaceBType are different types
	require.True(t, ptrIfaceAType != ptrIfaceBType)
	IfaceAType := ptrIfaceAType.Elem()
	IfaceBType := ptrIfaceBType.Elem()
	// even though InterfaceA and InterfaceB have the same method, they are not the same type
	require.True(t, IfaceAType != IfaceBType)

	var a InterfaceA
	require.Nil(t, a)
	require.Nil(t, reflect.TypeOf(a))
	a = &ImplA{}
	require.NotNil(t, a)

	// 'a' is a pointer to ImplA, so its reflect.Kind() is reflect.Ptr
	aType := reflect.TypeOf(a)
	require.Equal(t, reflect.Ptr, aType.Kind())
	require.Equal(t, "", aType.Name())
	require.Equal(t, "*reflect.ImplA", aType.String())
	require.False(t, aType.AssignableTo(ptrIfaceAType))
	// object of type *ImplA is assignable to InterfaceA
	require.True(t, aType.AssignableTo(IfaceAType))
	require.False(t, aType.AssignableTo(ptrIfaceAType))

	aElemType := aType.Elem()
	// object of type ImplA is assignable to InterfaceA
	require.False(t, aElemType.AssignableTo(ptrIfaceAType))
	require.True(t, aElemType.AssignableTo(IfaceAType))
	// object of type ImplA is not assignable to InterfaceB
	require.False(t, aElemType.AssignableTo(ptrIfaceBType))
	require.True(t, aElemType.AssignableTo(IfaceBType))
}
