package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	chan1, chan2 := make(chan int), make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		s, cnt1 := time.Now().Second(), 0
		for time.Now().Second()-s <= 1 {
			chan2 <- 2
			<-chan1
			cnt1++
		}
		fmt.Println(cnt1)
		wg.Done()
	}()
	go func() {
		s, cnt2 := time.Now().Second(), 0
		for time.Now().Second()-s <= 1 {
			<-chan2
			chan1 <- 1
			cnt2++
		}
		fmt.Println(cnt2)
		wg.Done()
	}()
	wg.Wait()
}
