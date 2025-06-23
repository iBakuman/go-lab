package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRangeNil(t *testing.T) {
	var s []int
	require.True(t, s == nil)
	require.NotPanics(t, func() {
		for i := range s {
			fmt.Println(i)
		}
	})
}