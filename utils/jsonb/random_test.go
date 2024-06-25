package jsonb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"go-lab/utils"
)

func TestGenRandomJson(t *testing.T) {
	requiredJsonPath := "root.all[0].email"
	jb, err := GenRandomJson(requiredJsonPath, 4)
	require.NoError(t, err)
	bs, err := jb.Marshal(context.Background())
	require.NoError(t, err)
	require.True(t, gjson.GetBytes(bs, "root.all.0.email").Exists())
	t.Logf(utils.MustBeautifyJson(string(bs)))
}
