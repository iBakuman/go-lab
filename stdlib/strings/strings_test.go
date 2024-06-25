package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFieldsFunc(t *testing.T) {
	str := "field1.field2[0].field3"
	fields := strings.FieldsFunc(str, func(r rune) bool {
		return r == '.' || r == '[' || r == ']'
	})
	require.Len(t, fields, 4)
	require.Equal(t, "field1", fields[0])
	require.Equal(t, "field2", fields[1])
	require.Equal(t, "0", fields[2])
	require.Equal(t, "field3", fields[3])
}
