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
