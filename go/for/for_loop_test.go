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
