// 可重入锁
package concurrent

import (
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

// 获取当前goroutine的id
func goId() int64 {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false) // false 不打印父goroutine的信息
	stackInfo := string(buf[:n])
	idstr := strings.Fields(strings.TrimPrefix(stackInfo, "goroutine "))[0]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		panic("Cannot get goroutine id")
	}
	return int64(id)
}

type ReentrantLock struct {
	mu        sync.Mutex
	owner     int64
	recursion int
}

func NewReentrantLock() *ReentrantLock {
	return &ReentrantLock{}
}

func (rl *ReentrantLock) Lock() {
	id := goId()
	if atomic.LoadInt64(&rl.owner) == id {
		rl.recursion++
		return
	}
	rl.mu.Lock()
	atomic.StoreInt64(&rl.owner, id)
	rl.recursion = 1
}

func (rl *ReentrantLock) Unlock() {
	id := goId()
	if atomic.LoadInt64(&rl.owner) != id {
		panic("Try unlock other goroutine")
	}
	rl.recursion--
	if rl.recursion != 0 {
		return
	}
	atomic.StoreInt64(&rl.owner, -1)
	rl.mu.Unlock()
}
