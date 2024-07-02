package unicode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUTF8(t *testing.T) {
	// \u8fc7
	a := '过'
	aStr := string(a)
	require.Equal(t, "[11101000 10111111 10000111]", fmt.Sprintf("%b", []byte(aStr)))
	t.Logf("%b\n", []byte(aStr))
	// 11101000 10111111 10000111
	// 1110xxxx 10xxxxxx 10xxxxxx
	//     1000   111111   000111
	//	   1000 1111 1100 0111
	//       8	 f     c    7
	t.Logf("%b\n", a)
	fmt.Println(reflect.TypeOf(a).Name())
}
