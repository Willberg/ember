package rwlock

import "sync"

type RWLock interface {
	Lock()
	UnLock()
	RLock()
	RUnLock()
}

// 两个锁实现读写锁
type MyRWLock struct {
	c    int
	r, g sync.Mutex
}

func (rw *MyRWLock) Lock() {
	rw.g.Lock()
}

func (rw *MyRWLock) UnLock() {
	rw.g.Unlock()
}

func (rw *MyRWLock) RLock() {
	rw.r.Lock()
	defer rw.r.Unlock()
	rw.c++
	if rw.c == 1 {
		rw.g.Lock()
	}
}

func (rw *MyRWLock) RUnLock() {
	rw.r.Lock()
	defer rw.r.Unlock()
	if rw.c == 0 {
		return
	}
	rw.c--
	if rw.c == 0 {
		rw.g.Unlock()
	}
}
