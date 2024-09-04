package channel_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// In Go, channels are comparable because they represent references to resources (communication endpoints) rather than
// the actual data being communicated. This comparability of channels allows Go to determine whether two channels are
// referring to the same underlying channel, which can be useful for certain operations and constructs.
//
// Key Points on Why Channels Are Comparable:
//
// 1. Channels as References:
//
// A channel in Go is a reference type, similar to pointers, slices, and maps. It represents a handle to a communication
// mechanism. When you compare two channels, you are essentially comparing their references (memory addresses), not the
// data being transmitted through them.
//
// 2. Comparability Based on Identity:
//
// Since channels are references, comparing them is straightforward. Two channels are considered equal if and only if
// they are the exact same channel (i.e., they have the same memory address). This comparability is like comparing two
// pointers or two interface values to see if they point to the same object.
//
// 3. Use Cases:
//
// Comparability is useful for determining if two variables are referring to the same channel. For example: In a select
// statement, you might want to check if different cases are using the same channel. When writing tests, you might want
// to assert that a function returns a specific channel. It can also be useful in complex scenarios like managing
// multiple channels, ensuring no duplicate channels in a slice, or other similar constructs.
//
// 4. Channel Equality and Zero Value:
//
// Comparing a channel with nil is a common pattern to check if the channel has been initialized. For example, if ch ==
// nil { ... } can be used to check if a channel has been properly set up. This check is crucial because operations on a
// nil channel block indefinitely, which can lead to deadlocks if not handled properly.

func TestComparability(t *testing.T) {
		ch1 := make(chan int)
		ch2 := make(chan int)
		ch3 := ch1

		require.False(t, ch1 == ch2) // false, different channels
		require.True(t, ch1 == ch3) // true, same channel

		ch1 = nil
		require.True(t, ch1 == nil) // true, channel is nil

		var ch4 chan int
		// not initialized, so it's nil
		require.True(t, ch4 == nil) // true, channel is nil
}

// Channels in Go are comparable because they represent references to communication endpoints. Comparability allows
// developers to determine if two channels refer to the same resource, which can be useful in various control and
// synchronization scenarios. This is consistent with Go's design philosophy of making low-level constructs, like
// references and channels, simple and predictable to use.