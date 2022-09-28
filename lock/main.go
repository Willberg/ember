package main

import (
	. "ember/lock/rwlock"
	"fmt"
	"sync"
)

func main() {
	rw := MyRWLock{}
	v, n, c := 0, 100, 10000
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(gid int) {
			for i := 0; i < c; i++ {
				rw.Lock()
				v++
				rw.UnLock()
				rw.RLock()
				fmt.Printf("%d, v = %d\n", gid, v)
				rw.RUnLock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
