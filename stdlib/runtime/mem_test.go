package runtime_test

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// Option 和 Options 类型定义
type Option interface {
	filter(s *state, t reflect.Type, vx, vy reflect.Value) applicableOption
}

type Options []Option

// state 和 applicableOption 示例定义
type state struct {
	curPath string
}

type applicableOption interface{}

// dummyOption 是一个实现 Option 的简单结构体
type dummyOption struct{}

func (d dummyOption) filter(s *state, t reflect.Type, vx, vy reflect.Value) applicableOption {
	return nil
}

// apply 方法实现
func (opts Options) apply(s *state, _, _ reflect.Value) {
	const warning = "ambiguous set of applicable options"
	const help = "consider using filters to ensure at most one Comparer or Transformer may apply"
	var ss []string
	for _, opt := range flattenOptions(nil, opts) {
		ss = append(ss, fmt.Sprint(opt))
	}
	set := strings.Join(ss, "\n\t")
	panic(fmt.Sprintf("%s at %#v:\n\t%s\n%s", warning, s.curPath, set, help))
}

// flattenOptions 示例实现
func flattenOptions(acc []Option, opts Options) []Option {
	for _, opt := range opts {
		acc = append(acc, opt)
	}
	return acc
}

func TestMallocs(t *testing.T) {
	// 创建一个 Options 和 state 实例
	opts := Options{dummyOption{}, dummyOption{}}
	s := &state{curPath: "root"}

	// 使用 runtime.MemStats 统计内存分配情况
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	allocsBefore := m.Mallocs

	// 多次调用 apply 方法
	for i := 0; i < 1000; i++ {
		func() {
			defer func() {
				// 忽略 panic，只为了测试
				recover()
			}()
			opts.apply(s, reflect.Value{}, reflect.Value{})
		}()
	}

	runtime.ReadMemStats(&m)
	allocsAfter := m.Mallocs

	// 输出内存分配次数变化
	fmt.Printf("Memory allocations before: %d\n", allocsBefore)
	fmt.Printf("Memory allocations after: %d\n", allocsAfter)
	fmt.Printf("Allocations difference: %d\n", allocsAfter-allocsBefore)
}
