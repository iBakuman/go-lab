package go_cmp_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func equateIfZero(t *testing.T, x, y interface{}) string {
	alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
	// This option handles slices and maps of any type.
	opt := cmp.FilterValues(func(x, y interface{}) bool {
		vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
		return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
			(vx.IsZero() || vy.IsZero())
	}, alwaysEqual)
	diff := cmp.Diff(x, y, opt)
	t.Logf(diff)
	return diff
}

type Human struct {
	Name    string
	Age     int
	Scores  map[string]int
	Friends []*Human
}

func Test_equateIfZero(t *testing.T) {
	h1 := Human{
		Name: "Alice",
		Age:  20,
		Scores: map[string]int{
			"Math": 90,
			"Eng":  80,
			"Art":  70,
		},
		Friends: []*Human{
			{
				Name: "Bob",
				Age:  21,
				Scores: map[string]int{
					"Math": 85,
				},
			},
		},
	}
	h2 := Human{
		Name: "Alice",
		Age:  20,
		Scores: map[string]int{
			"Math": 90,
			"Eng":  80,
			"Art":  70,
		},
		Friends: []*Human{
			{
				Name: "Bob",
				Age:  21,
				Scores: map[string]int{
					"Math": 85,
				},
			},
		},
	}

	t.Run("equal", func(t *testing.T) {
		require.Empty(t, equateIfZero(t, h1, h2))
	})

	t.Run("different age", func(t *testing.T) {
		oldValue := h2.Age
		t.Cleanup(func() {
			h2.Age = oldValue
		})
		h2.Age = 21
		require.NotEmpty(t, equateIfZero(t, h1, h2))
	})

	t.Run("nil friends", func(t *testing.T) {
		oldValue := h2.Friends
		t.Cleanup(func() {
			h2.Friends = oldValue
		})
		h2.Friends = nil
		require.Empty(t, equateIfZero(t, h1, h2))
	})

	t.Run("different friend", func(t *testing.T) {
		oldValues := h2.Friends
		t.Cleanup(func() {
			h2.Friends = oldValues
		})
		h2.Friends[0].Age = 22
		require.NotEmpty(t, equateIfZero(t, h1, h2))
	})

	t.Run("nil scores", func(t *testing.T) {
		oldValue := h2.Scores
		t.Cleanup(func() {
			h2.Scores = oldValue
		})
		h2.Scores = nil
		require.Empty(t, equateIfZero(t, h1, h2))
	})

	t.Run("empty scores", func(t *testing.T) {
		oldValue := h2.Scores
		t.Cleanup(func() {
			h2.Scores = oldValue
		})
		h2.Scores = map[string]int{}
		// Note: The map is not zero value, so it should not be equal.
		require.NotEmpty(t, equateIfZero(t, h1, h2))
	})
}
