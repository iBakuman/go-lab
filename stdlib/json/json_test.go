package json

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"go-lab/utils/jsonb"
)

func TestDecoder(t *testing.T) {
	arrElems := []any{"one", "two", "three"}
	jsonArr := jsonb.A(arrElems...)
	jsonArrStr, err := jsonArr.Marshal(context.Background())
	require.NoError(t, err)
	decoder := json.NewDecoder(bytes.NewReader(jsonArrStr))
	token, err := decoder.Token()
	require.NoError(t, err)
	require.EqualValues(t, json.Delim('['), token)
	for _, elem := range arrElems {
		var s string
		err := decoder.Decode(&s)
		require.NoError(t, err)
		require.EqualValues(t, elem, s)
	}
	token, err = decoder.Token()
	require.NoError(t, err)
	require.EqualValues(t, json.Delim(']'), token)
}
