package context

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func TestValue(t *testing.T) {
	type contextKey int
	const (
		key1 contextKey = 1
		key2            = 2
		key3            = 3
	)
	newContext := func(ctx context.Context, k contextKey, v any, convert bool) context.Context {
		if convert {
			return context.WithValue(ctx, int(k), v)
		} else {
			return context.WithValue(ctx, k, v)
		}
	}
	fromContext := func(ctx context.Context, k contextKey, convert bool) any {
		if convert {
			return ctx.Value(int(k))
		} else {
			return ctx.Value(k)
		}
	}

	type person struct {
		Name string
		Age  int
	}

	t.Run("plain", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "key", "value")
		require.Equal(t, "value", ctx.Value("key"))
		require.Nil(t, ctx.Value("not-exist"))
		require.Panics(t, func() {
			_ = ctx.Value("non-exist").(*person)
		})
		// if the key is not exist, the value will be nil
		v, ok := ctx.Value("non-exist").(*person)
		require.Nil(t, v)
		require.False(t, ok)
	})

	t.Run("stored in contextKey, retrieved with contextKey", func(t *testing.T) {
		ctx := context.Background()
		ctx = newContext(ctx, key1, &person{Name: "Alice", Age: 20}, false)
		require.Equal(t, "Alice", fromContext(ctx, key1, false).(*person).Name)
	})

	t.Run("stored in int, retrieved with int", func(t *testing.T) {
		ctx := context.Background()
		ctx = newContext(ctx, key1, &person{Name: "Alice", Age: 20}, true)
		require.Equal(t, "Alice", fromContext(ctx, key1, true).(*person).Name)
	})

	t.Run("stored in contextKey, retrieved with int", func(t *testing.T) {
		ctx := context.Background()
		ctx = newContext(ctx, key1, &person{Name: "Alice", Age: 20}, false)
		// type contextKey not eqaul to int, so the value will be nil
		require.Nil(t, fromContext(ctx, key1, true))
		require.Equal(t, "contextKey", reflect.TypeOf(key1).Name())
		require.Equal(t, "int", reflect.TypeOf(int(key1)).Name())
		require.Equal(t, "int", reflect.ValueOf(key1).Type().Name())
	})

	t.Run("stored in int, retrieved with contextKey", func(t *testing.T) {
		ctx := context.Background()
		ctx = newContext(ctx, key1, &person{Name: "Alice", Age: 20}, true)
		require.Nil(t, fromContext(ctx, key1, false))
	})
}

func TestValueWithEmptyContext(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}
	ctx := context.WithValue(context.Background(), "key", "value")
	require.Equal(t, "value", ctx.Value("key"))
	require.Nil(t, ctx.Value("not-exist"))
	require.Panics(t, func() {
		_ = ctx.Value("non-exist").(*person)
	})
	v, ok := ctx.Value("non-exist").(*person)
	require.Nil(t, v)
	require.False(t, ok)

	type e1 struct{}
	a := e1{}
	b := e1{}
	require.True(t, a == b)
	type e3 struct{}
	a1 := e3{}
	var ia any = a
	var ia1 any = a1
	require.False(t, ia == ia1)

	type e2 struct {
		Name string
	}
	c := e2{Name: "Alice"}
	d := e2{Name: "Alice"}
	require.True(t, c == d)
	var i1 any = c
	var i2 any = d
	require.True(t, i1 == i2)

	c = e2{Name: "Alice"}
	d = e2{Name: "Bob"}
	require.False(t, c == d)
	i1 = c
	i2 = d
	require.False(t, i1 == i2)
}

func TestContext(t *testing.T) {
	ctx := context.Background()
	funA := func(ctx context.Context) {
		ctx = context.WithValue(ctx, "funA", "funA")
		require.NotNil(t, ctx.Value("funA"))
	}
	funA(ctx)
	require.Nil(t, ctx.Value("funA"))
}
