package time_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go-lab/utils"
)

const (
	iterations = 1000000
)

func TestTimeAfterMemoryLeak(t *testing.T) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	initialAlloc := memStats.Alloc
	for i := 0; i < iterations; i++ {
		go func() {
			time.After(10 * time.Second)
		}()
	}

	// wait for goroutines to start
	time.Sleep(1 * time.Second)
	runtime.ReadMemStats(&memStats)
	finalAlloc := memStats.Alloc

	fmt.Printf("Initial memory allocation: %.2f MB\n", utils.BytesToMB(initialAlloc))
	fmt.Printf("Final memory allocation: %.2f MB\n", utils.BytesToMB(finalAlloc))
	fmt.Printf("Memory increased by: %.2f MB\n", utils.BytesToMB(finalAlloc-initialAlloc))
	assert.Greater(t, int(finalAlloc-initialAlloc), 10*1024*1024)
}

func TestTimeTimerNoMemoryLeak(t *testing.T) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	initialAlloc := memStats.Alloc
	for i := 0; i < iterations; i++ {
		go func() {
			timer := time.NewTimer(10 * time.Second)
			timer.Stop()
			<-timer.C
		}()
	}
	time.Sleep(1 * time.Second)
	runtime.GC()
	runtime.ReadMemStats(&memStats)
	finalAlloc := memStats.Alloc

	fmt.Printf("Initial memory allocation: %.2f MB\n", utils.BytesToMB(initialAlloc))
	fmt.Printf("Final memory allocation: %.2f MB\n", utils.BytesToMB(finalAlloc))
	fmt.Printf("Memory increased by: %.2f MB\n", utils.BytesToMB(finalAlloc-initialAlloc))

	assert.Less(t, int(finalAlloc-initialAlloc), 10*1024*1024)
}

func TestExecuteTaskPeriodically(t *testing.T) {
	// execute task periodically
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			executeTask(t)
		}
	}
}

func executeTask(t *testing.T) {
	fmt.Println("Executing task at", time.Now())
	// 在这里执行你的任务
	time.Sleep(2 * time.Second) // 模拟任务执行时间
}
