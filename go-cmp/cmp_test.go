package go_cmp

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

type A struct {
	Name            string
	Age             int
	Time            time.Time
	Field1          string
	Map             map[string]string
	unexportedField string
}

func GenA() A {
	return A{
		Name:            "A",
		Age:             10,
		Field1:          "Field1",
		Map:             map[string]string{"key1": "value1"},
		unexportedField: "unexportedField",
	}
}

func TestCmp(t *testing.T) {
	a := GenA()
	b := GenA()
	require.Panics(t, func() {
		cmp.Equal(a, b)
	})
	require.NotPanics(t, func() {
		require.False(t, cmp.Equal(a, b), cmpopts.IgnoreUnexported(a))
		require.True(t, cmp.Equal(a, b, cmpopts.IgnoreUnexported(a, time.Time{})))
	})
}

type B struct {
	A    A
	Type string
}

func TestCmpB(t *testing.T) {
	a1 := GenA()
	a2 := GenA()
}

func TestWhile(t *testing.T) {
	for {
		t.Logf("while loop")
		time.Sleep(time.Second)
	}
}
