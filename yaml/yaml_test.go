package yaml_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestUnmarshal(t *testing.T) {
	bs := []byte("a: 1\nb: 20m")
	type ty struct {
		A int `yaml:"a"`
		B time.Duration
	}
	var i ty
	require.NoError(t, yaml.Unmarshal(bs, &i))
	require.Equal(t, 1, i.A)
	require.Equal(t, 20*time.Minute, i.B)
}
