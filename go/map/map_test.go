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

func TestCreate(t *testing.T) {
	var m1 map[string]string
	require.Nil(t, m1)
	require.Empty(t, m1["key"])
	require.Len(t, m1, 0)

	m2 := map[string]string{}
	require.NotNil(t, m2)
	// we only access the key, so it should not create the key in the map
	require.Empty(t, m2["key"])
	require.Len(t, m2, 0)
	m2["key"] = "value"
	require.Equal(t, "value", m2["key"])
	require.Len(t, m2, 1)
}
