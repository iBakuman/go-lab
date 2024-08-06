package map_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNIl(t *testing.T) {
	var m1 map[string]string
	require.Nil(t, m1)
	require.Empty(t, m1["key"])
	require.Panics(t, func() {
		m1["key"] = "value"
	})
	m2 := map[string]string{}
	require.NotNil(t, m2)
	require.Empty(t, m2["key"])
	require.NotPanics(t, func() {
		m2["key"] = "value"
	})
	require.Equal(t, "value", m2["key"])
}
