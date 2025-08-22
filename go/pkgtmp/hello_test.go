package pkgtmp

import (
	"context"
	"github.com/ibakuman/go-lab/go/pkgtmp/testkit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSayHello(t *testing.T) {
	require.NoError(t, SayHello(context.Background(), testkit.NewHandler()))
}
