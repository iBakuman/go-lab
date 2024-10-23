package reflect

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTypeOf(t *testing.T) {

	t.Run("get type of a variable", func(t *testing.T) {
		// get the type of time.Time
		timeType := reflect.TypeOf(time.Time{})
		// also works, and is more efficient, because it doesn't create a new instance of time.Time
		timeType = reflect.TypeOf((*time.Time)(nil)).Elem()

		var a interface{}
		// before assigning a value to a, its type is nil
		require.True(t, reflect.TypeOf(a) == nil)
		a = time.Now()
		// after assigning a value to a, its type is time.Time
		require.True(t, reflect.TypeOf(a) == timeType)
		// a is a time.Time, so this type assertion will not panic
		require.True(t, reflect.TypeOf(a.(time.Time)) == timeType)

	})

	t.Run("empty interface", func(t *testing.T) {
		// myInterface is an empty interface, which can hold values of any type.
		// reflect.TypeOf returns the dynamic type of the interface value c.
		type myInterface interface{}
		var c myInterface
		require.True(t, reflect.TypeOf(c) == nil)
		require.Panics(t, func() {
			// reflect.TypeOf(nil) will panic.
			reflect.TypeOf(c).Kind()
		})
		c = 0
		// this call will not panic
		require.True(t, reflect.TypeOf(c).Kind() == reflect.Int)
	})

	t.Run("pointer to interface", func(t *testing.T) {
		type ifaceA interface {
			foo()
		}
		type ifaceB interface {
			foo()
		}
		var pA *ifaceA
		var pB *ifaceB
		require.True(t, pA == nil)
		require.True(t, pB == nil)
		// compile error: cannot compare pA and pB because they are of different types
		// require.True(t, pA != pB)
		// even though pA and pB are both nil, their types are both non-nil
		// because they are pointers to interfaces, not interfaces.
		require.True(t, reflect.TypeOf(pA) != nil)
		require.True(t, reflect.TypeOf(pB) != nil)

		var pA2 ifaceA
		var pB2 ifaceB
		// pA2 and pB2 are both nil, and their types are nil as well.
		require.True(t, reflect.TypeOf(pA2) == nil)
		require.True(t, reflect.TypeOf(pB2) == nil)
	})
}

type Person struct {
	name     string
	padding1 [1000]int
	padding2 [1000]int
}

var (
	globalSliceA []*Person
	globalSliceB []*any
)

func FuncA(person Person) {
	globalSliceA = append(globalSliceA, &person)
}

func FuncB(person any) {
	globalSliceB = append(globalSliceB, &person)
}

func FuncC(person any) {
}

func memoryConsumed(f func()) uint64 {
	var m1, m2 runtime.MemStats
	runtime.GC()               // 运行垃圾回收，确保所有内存都被释放
	runtime.ReadMemStats(&m1)  // 读取当前内存使用情况
	f()                        // 执行测试函数
	runtime.ReadMemStats(&m2)  // 读取函数执行后的内存使用情况
	return m2.Alloc - m1.Alloc // 返回内存增量
}

func TestMemoryUsage(t *testing.T) {
	p := Person{name: "Alice"}

	// 测试 FuncA 的内存使用
	memA := memoryConsumed(func() {
		for i := 0; i < 1000000; i++ {
			FuncA(p)
		}
	})

	// 测试 FuncB 的内存使用
	memB := memoryConsumed(func() {
		for i := 0; i < 1000000; i++ {
			FuncB(p)
		}
	})

	fmt.Printf("FuncA memory usage: %d bytes\n", memA)
	fmt.Printf("FuncB memory usage: %d bytes\n", memB)
}
