package sort

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	t.Run("sort int slice", func(t *testing.T) {
		elems := []int{5, 3, 4, 1, 2}
		sort.Slice(elems, func(i, j int) bool {
			return elems[i] < elems[j]
		})
		for i := 1; i < len(elems); i++ {
			require.LessOrEqual(t, elems[i-1], elems[i])
		}
		sort.Slice(elems, func(i, j int) bool {
			return elems[i] > elems[j]
		})
		for i := 1; i < len(elems); i++ {
			require.GreaterOrEqual(t, elems[i-1], elems[i])
		}
	})
	t.Run("sort time slice", func(t *testing.T) {
		elems := []time.Time{
			time.Now().Add(-time.Hour),
			time.Now().Add(-time.Minute),
			time.Now().Add(-time.Second),
			time.Time{},
			time.Now(),
		}
		sort.Slice(elems, func(i, j int) bool {
			if elems[i].IsZero() {
				return false
			}
			if elems[j].IsZero() {
				return true
			}
			return elems[i].Before(elems[j])
		})
		for i := 1; i < len(elems)-1; i++ {
			require.True(t, elems[i-1].Before(elems[i]))
		}
		require.True(t, elems[len(elems)-1].IsZero())
		sort.Slice(elems, func(i, j int) bool {
			if elems[i].IsZero() {
				return true
			}
			if elems[j].IsZero() {
				return false
			}
			return elems[i].After(elems[j])
		})
		for i := 1; i < len(elems)-1; i++ {
			require.True(t, elems[i].After(elems[i+1]))
		}
		require.True(t, elems[0].IsZero())
	})
}
