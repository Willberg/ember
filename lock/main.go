package main

import (
	. "ember/lock/rwlock"
	"fmt"
	"sync"
)

func main() {
	rw := MyRWLock{}
	v, n := 0, 5
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(gid int) {
			for i := 0; i < 10; i++ {
				rw.RLock()
				fmt.Printf("%d, v = %d\n", gid, v)
				rw.RUnLock()
				rw.Lock()
				cur := v
				cur++
				v = cur
				rw.UnLock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
