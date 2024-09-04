package _for

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForLoop(t *testing.T) {
	type Animal struct {
		Name string `custom:"a"`
		Age  int
	}
	ty := reflect.TypeOf((*Animal)(nil)).Elem()
	require.Equal(t, "Animal", ty.Name())
	require.Equal(t, "Name", ty.Field(0).Name)
	require.Equal(t, "a", ty.Field(0).Tag.Get("custom"))

	ty2 := reflect.TypeOf(&Animal{})
	require.Empty(t, ty2.Name())
	require.Equal(t, reflect.Ptr, ty2.Kind())

	animals := []Animal{
		{Name: "dog", Age: 1},
		{Name: "cat", Age: 2},
	}

	var d1 []Animal
	for _, a := range animals {
		ty3 := reflect.TypeOf(a)
		require.Equal(t, ty, ty3)
		a.Name = "bird"
		d1 = append(d1, a)
	}

	// make sure the `a.Name = "bird"` line above doesn't affect the original animals
	for _, a := range animals {
		t.Log(a.Name)
		require.NotEqual(t, "bird", a.Name)
	}
	// &d1[0] and &d1[1] point to different memory addresses.
	require.NotEqual(t, fmt.Sprintf("%p", &d1[0]), fmt.Sprintf("%p", &d1[1]))
	// not equal because the age is different.
	require.NotEqual(t, &d1[0], &d1[1])

	var d2 []Animal
	dog := Animal{Name: "dog", Age: 1}
	d2 = append(d2, dog)
	d2 = append(d2, dog)
	require.NotEqual(t, fmt.Sprintf("%p", &d2[0]), fmt.Sprintf("%p", &d2[1]))
	// Amazingly, &d2[0] and &d2[1] are equal, even though they point to different memory addresses.
	// Seems like require.Equal compares the values of the two pointers, not the memory addresses.
	require.Equal(t, &d2[0], &d2[1])
}

func TestTripleForLoop(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	for i := 0; i < 5; i++ {
		t.Logf("--------------------------------------------------------")
		t.Logf("i: %d", i)
		for j := 0; j < 3; j++ {
			t.Logf("+++++++++++++++++++++++++++++++++++++++")
			t.Logf("j: %d", j)
			select {
			case <-ch:
				for k := 0; k < 4; k++ {
					t.Logf("==================")
					t.Logf("k: %d", k)
					if k == 1 {
						break
					}
				}
			}
			if j == 2 {
				break
			}
		}
		if i == 3 {
			break
		}
	}
}

func TestRange(t *testing.T) {
	type Animal struct {
		name string
		legs int
	}

	zoo := []Animal{
		{"Dog", 4},
		{"Chicken", 2},
		{"Snail", 0},
	}

	fmt.Printf("-> Before update %v\n", zoo)
	for _, animal := range zoo {
		// 🚨 Oppps! `animal` is a copy of an element 😧
		animal.legs = 999
	}
	fmt.Printf("\n-> After update %v\n", zoo)
	for i := range zoo {
		t.Logf("idx: %d, ptr: %p", i, &i)
	}
}

// Ranging Over Nil: If you attempt to range over a nil slice, map, or channel,
// the `for range` loop will not execute and will not result in a runtime panic.
// However, ranging over a nil pointer to an array will result in a runtime panic.
// Always check if a pointer to an array is nil before ranging over it.
func TestRangeNilSlice(t *testing.T) {
	var slice []int
	for range slice {
		require.Fail(t, "should not reach here")
	}
	var m map[int]int
	for range m {
		require.Fail(t, "should not reach here")
	}
	var ptrSlice = &[]int{}
	for range *ptrSlice {
		require.Fail(t, "should not reach here")
	}

	var ptrArray *[3]int
	var cnt int
	require.NotPanics(t, func() {
		// we can range over a nil pointer to an array, but we cannot get the value.(only the index)
		for idx := range ptrArray {
			_ = idx
			cnt++
			// require.Equal(t, 0, v)
		}
	})
	require.Equal(t, 3, cnt)
	require.Panics(t, func() {
		// if we attempt to range over a nil pointer to an array and get the value,
		// it will result in a runtime panic.
		for _, v := range *ptrArray {
			_ = v
		}
	})
}
