package reflect

import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type Printer interface{
	Print() string
}

type PrinterImpl struct{
	msg string
}

func (p PrinterImpl) Print() string {
	return p.msg
}

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

		type sA struct {
			name string
		}
		var psA *sA
		var b interface{}
		b = psA
		require.True(t, b != nil)
		require.True(t, reflect.TypeOf(b) == reflect.TypeOf((*sA)(nil)))
		require.True(t, reflect.TypeOf(b).Elem() == reflect.TypeOf(sA{}))
		// even though psA is nil, reflect.TypeOf(a) is not nil.
		require.True(t, reflect.TypeOf(b) != nil)
		var vSA sA
		b = vSA
		require.True(t, reflect.TypeOf(b) == reflect.TypeOf(sA{}))
	})

	t.Run("empty interface", func(t *testing.T) {
		// myInterface is an empty interface, which can hold values of any type.
		// reflect.TypeOf returns the dynamic type of the interface value c.
		type myInterface interface{}
		var c myInterface
		require.True(t, reflect.TypeOf(c) == nil)
		require.Panics(t, func() {
			// ⚠️Attention⚠️: this will panic because c is nil
			reflect.TypeOf(c).Kind()
		})
		c = 0
		// this call will not panic
		require.True(t, reflect.TypeOf(c).Kind() == reflect.Int)
	})

	t.Run("non-empty interface", func(t *testing.T) {
		type myInterface interface {
			foo()
		}
		var c myInterface
		// ⚠️Attention⚠️: even though myInterface has a method foo, c is nil.
		require.True(t, reflect.TypeOf(c) == nil)
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
		// because they are pointers to interfaces rather than interfaces themselves.
		require.True(t, reflect.TypeOf(pA) != nil)
		require.True(t, reflect.TypeOf(pB) != nil)

		var pA2 ifaceA
		var pB2 ifaceB
		require.True(t, reflect.TypeOf(pA2) == nil)
		require.True(t, reflect.TypeOf(pB2) == nil)
	})
	t.Run("interface kind", func(t *testing.T) {
		var p Printer = PrinterImpl{msg: "hello"}
		require.True(t, reflect.TypeOf(p).Kind() == reflect.Struct)
		a := make(map[string]interface{})
		a["key"] = p
		require.True(t, reflect.TypeOf(a["key"]).Kind() == reflect.Struct)
		var c interface{}
		c = io.Closer(nil)
		require.True(t, c == nil)
		require.True(t, reflect.TypeOf(c) == nil)
		require.True(t, reflect.ValueOf(c).IsValid() == false)
		require.True(t, reflect.ValueOf(c) == reflect.Value{})
	})

	t.Run("reflect.Interface", func(t *testing.T) {
		type Example struct{
			Field interface{}
		}
		e := Example{Field: "Hello"}
		v := reflect.ValueOf(e)
		require.True(t, v.FieldByName("Field").Kind() == reflect.Interface)
		var strA interface{} = "example"
		v = reflect.ValueOf(strA)
		require.True(t, v.Kind() == reflect.String)
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
