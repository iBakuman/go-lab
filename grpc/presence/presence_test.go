package presence

import (
	"testing"

	"github.com/ibakuman/go-lab/grpc/gen/presence"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFieldPresence(t *testing.T) {
	m1 := &presence.MessageB{}
	data, err := proto.Marshal(m1)
	require.NoError(t, err)
	// mA has no field set, so it should be empty(i.e len(data) == 0)
	require.Empty(t, data)
	require.Len(t, data, 0)

	m2 := &presence.MessageB{}
	err = proto.Unmarshal(data, m2)
	require.NoError(t, err)
	// message-type fields support field presence, so the field should be nil
	require.True(t, m2.B == nil)
	// scalar-type fields do not support field presence, so the field should be set to the zero value.
	// we can't distinguish between a field that was set to the zero value and a field that was not set at all.
	require.Equal(t, m2.A, int32(0))
	// optional scalar-type support field presence, so the field should be nil.
	require.Equal(t, m2.C, (*int32)(nil))
	require.False(t, m2.ProtoReflect().Has(m2.ProtoReflect().Descriptor().Fields().ByName("c")))

	val := int32(0)
	m3 := &presence.MessageB{A: 0, C: &val}
	data, err = proto.Marshal(m3)
	require.NoError(t, err)
	m4 := &presence.MessageB{}
	err = proto.Unmarshal(data, m4)
	require.NoError(t, err)
	// we can't distinguish between a field that was set to the zero value and a field that was not set at all.
	require.Equal(t, m4.A, int32(0))
	// optional scalar-type support field presence, so the field should be nil.
	require.NotNil(t, m4.C)
	require.Equal(t, *m4.C, int32(0))
	require.True(t, m4.ProtoReflect().Has(m4.ProtoReflect().Descriptor().Fields().ByName("c")))
}
