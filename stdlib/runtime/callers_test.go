package runtime

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func A(skip int) string {
	return B(skip)
}

func B(skip int) string {
	pc := make([]uintptr, 10)
	runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc)
	var sb strings.Builder
	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf("%s:%d\n", frame.Function, frame.Line))
		if !more {
			break
		}
	}
	return sb.String()
}

func TestCaller(t *testing.T) {
	s1 := A(0)
	require.Contains(t, s1, "runtime.Callers")
	require.Contains(t, s1, "runtime.B")
	require.Contains(t, s1, "runtime.A")

	s2 := A(1)
	require.NotContains(t, s2, "runtime.Callers")
	require.Contains(t, s2, "runtime.B")
	require.Contains(t, s2, "runtime.A")

	s3 := A(2)
	require.NotContains(t, s3, "runtime.B")
	require.Contains(t, s3, "runtime.A")
}
