package nil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptySlice(t *testing.T) {
	var nilVar []int
	require.Nil(t, nilVar)
	nonNilVar := make([]int, 0)
	require.NotNil(t, nonNilVar)
}
