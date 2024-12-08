package pointer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSinglePointer(t *testing.T) {
	funcA := func(i *int) {
		*i = 10
	}
	i := 20
	funcA(&i)
	require.Equal(t, 10, i)
}

func TestDoublePointer(t *testing.T) {
	funcA := func(i **int) {
		**i = 10
	}
	i := 20
	b := &i
	funcA(&b)
	// NOTE: This is not allowed in Go
	// funcA(&(&i))
	require.Equal(t, 10, i)
}

func TestTriplePointer(t *testing.T) {
	funcA := func(i ***int) {
		//***i = 10
		a := 1
		b := &a
		c := &b
		*i = c
	}
	i := 20
	b := &i
	c := &b
	funcA(&c)
	require.Equal(t, 10, i)
}
