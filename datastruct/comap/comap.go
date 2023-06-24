package comap

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type MyConcurrentMap struct {
	sync.Mutex
	keyToCh map[int]chan struct{}
	mp      map[int]int
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		keyToCh: make(map[int]chan struct{}),
		mp:      make(map[int]int),
	}
}

func (m *MyConcurrentMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v

	ch, ok := m.keyToCh[k]
	if !ok {
		return
	}

	// 关闭chan可以唤醒所有阻塞的读go routine, 但是只能关闭一次
	select {
	case <-ch:
		fmt.Println("已经关闭一次ch")
		return
	default:
		fmt.Println("写入 关闭ch")
		close(ch)
	}
}

func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	v, ok := m.mp[k]
	if ok {
		m.Unlock()
		return v, nil
	}

	ch, ok := m.keyToCh[k]
	if !ok {
		ch = make(chan struct{})
		m.keyToCh[k] = ch
	}
	ctx, cancel := context.WithTimeout(context.Background(), maxWaitingDuration)
	defer cancel()
	m.Unlock()

	select {
	case <-ctx.Done():
		return -1, ctx.Err()
	case <-ch:
		fmt.Println("test 读 阶段1")
	}

	fmt.Println("test 读 阶段2")
	m.Lock()
	defer m.Unlock()
	return m.mp[k], nil
}
