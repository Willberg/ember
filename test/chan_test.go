package test

import (
	"ember/datastruct/comap"
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

func TestChan0(t *testing.T) {
	ch := make(chan int, 4)
	var wg sync.WaitGroup
	wg.Add(3)
	defer wg.Wait()
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
		wg.Done()
	}()
	for i := 0; i < 2; i++ {
		go func(p int) {
		loop:
			for {
				select {
				case x, ok := <-ch:
					// 通道被关闭之后，接受操作会一直接受对应类型的零值，发送操作会宕机
					if !ok {
						break loop
					}
					fmt.Printf("%d, %d\n", p, x)
					time.Sleep(500 * time.Millisecond)
				}
			}
			wg.Done()
		}(i)
	}

}

func TestChan1(t *testing.T) {
	pn, cn, pnn, chn := 5, 20, 4000, 10
	var pw, cw time.Duration = 10, 50
	ch := make(chan string, chn)
	var pwg, cwg sync.WaitGroup
	pwg.Add(pn)
	cwg.Add(cn)
	defer cwg.Wait()
	for i := 0; i < pn; i++ {
		go func(i int) {
			for j := 1; j <= pnn; j++ {
				ch <- fmt.Sprintf("pid:%d, v=%d, time: %s", i, j, time.Now().Format("2006-01-02 15:04:05"))
				time.Sleep(pw * time.Millisecond)
			}
			pwg.Done()
		}(i)
	}
	for i := 0; i < cn; i++ {
		go func(i int) {
		loop:
			for {
				select {
				case s, ok := <-ch:
					if !ok {
						break loop
					}
					fmt.Println(fmt.Sprintf("%s, c:%d, now: %s", s, i, time.Now().Format("2006-01-02 15:04:05")))
					time.Sleep(cw * time.Millisecond)
				}
			}
			cwg.Done()
		}(i)
	}
	pwg.Wait()
	close(ch)
}

func TestChan2(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 1
	v, ok := <-ch
	fmt.Println(v, ok)
	close(ch)
	v, ok = <-ch
	fmt.Println(v, ok)
}

func TestMyConcurrentChan1(t *testing.T) {
	cmap := comap.NewMyConcurrentMap()
	n := sync.WaitGroup{}
	n.Add(4)
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println("读取1")
			v, err := cmap.Get(1, 5*time.Second)
			fmt.Println("读取1 结束")
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			fmt.Println(v)
			n.Done()
		}()
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println("写入1")
			cmap.Put(1, 1)
			fmt.Println("写入1 结束")
			n.Done()
		}()
	}
	n.Wait()
}

func TestMyConcurrentChan2(t *testing.T) {
	cmap := comap.NewMyConcurrentMap2()
	n := sync.WaitGroup{}
	n.Add(30)
	for i := 0; i < 20; i++ {
		go func() {
			fmt.Println("读取1")
			v, err := cmap.Get(1, 5*time.Second)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			fmt.Println(v)
			fmt.Println("读取1 结束")
			n.Done()
		}()
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("写入1")
			cmap.Put(1, 1)
			fmt.Println("写入1 结束")
			n.Done()
		}()
	}
	n.Wait()
}
