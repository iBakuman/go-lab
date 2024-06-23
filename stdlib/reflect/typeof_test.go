package reflect

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTypeOf(t *testing.T) {
	timeType := reflect.TypeOf(time.Time{})
	var a interface{}
	require.True(t, reflect.TypeOf(a) == nil)
	a = time.Now()
	require.True(t, reflect.TypeOf(a) == timeType)
	require.True(t, reflect.TypeOf(a.(time.Time)) == timeType)
	var b interface{} = a
	require.True(t, reflect.TypeOf(b) == timeType)
	require.True(t, reflect.TypeOf(b.(time.Time)) == timeType)
	type myInterface interface{}
	var c myInterface
	require.True(t, reflect.TypeOf(c).Kind() == reflect.Interface)
}
