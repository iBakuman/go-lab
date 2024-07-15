package _defer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefer(t *testing.T) {
	fn := func() {
		return
		defer func() {
			t.Fatal("should not reach here")
		}()
	}
	fn()
}

func TestDefer2(t *testing.T) {
	t.Run("return value from recover should be string", func(t *testing.T) {
		defer func() {
			recovered := recover()
			ty := reflect.TypeOf(recovered)
			require.Equal(t, reflect.TypeOf(""), ty)
		}()
		err := "error"
		panic(err)
	})

	t.Run("return value from recover should be error", func(t *testing.T) {
		defer func() {
			recovered := recover()
			_, ok := recovered.(error)
			require.True(t, ok)
			ty := reflect.TypeOf(recovered)
			require.NotEqual(t, reflect.TypeOf((*error)(nil)).Elem(), ty)
			require.Equal(t, reflect.TypeOf(errors.New("")), ty)
		}()
		err := errors.New("error")
		panic(err)
	})
}

func TestDefer3(t *testing.T) {
	// recover be called in 3rd defer
	f1 := func() {
		r := recover()
		t.Logf("f2 is called, recover: %v", r)
	}

	f2 := func() {
		t.Logf("f1 is called")
		f1()
	}

	f3 := func() {
		t.Logf("f3 is called")
		defer f1()
	}

	require.Panics(t, func() {
		defer f2()
		panic("panic")
	})

	require.Panics(t, func() {
		defer f3()
		panic("panic")
	})

	require.NotPanics(t, func() {
		defer f1()
		panic("panic")
	})
}
