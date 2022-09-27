package main

import (
	. "ember/algorithm/loadbalance/robin"
	"fmt"
	"sync"
)

func normalSmoothRobinTest() {
	var pairs []Pair
	pairs = append(pairs, Pair{Ip: "192.168.0.1", Weight: 5})
	pairs = append(pairs, Pair{Ip: "192.168.0.2", Weight: 2})
	pairs = append(pairs, Pair{Ip: "192.168.0.3", Weight: 1})
	record := make(map[string][]int)
	for i, p := range pairs {
		record[p.Ip] = []int{p.Weight, i}
	}
	cnt := make(map[string]int)
	ps := CreateSmoothRobin(pairs)
	for i := 0; i < 16; i++ {
		ip := ps.GetPeer()
		cnt[ip]++
		fmt.Println(ip, record[ip][0], record[ip][1])
	}
	for k, v := range cnt {
		fmt.Println(k, v)
	}
	fmt.Printf("\n\n\n")
}

func createSmoothRobin(once *sync.Once, r *Robin) {
	once.Do(func() {
		var pairs []Pair
		pairs = append(pairs, Pair{Ip: "192.168.0.1", Weight: 5})
		pairs = append(pairs, Pair{Ip: "192.168.0.2", Weight: 2})
		pairs = append(pairs, Pair{Ip: "192.168.0.3", Weight: 1})
		*r = CreateSmoothRobin(pairs)
		fmt.Println("smooth robin init...")
	})
}

func parallelSmoothRobinTest() {
	var wg sync.WaitGroup
	var once sync.Once
	var r Robin
	n := 5
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(pid int, once *sync.Once, r *Robin) {
			createSmoothRobin(once, r)
			for i := 0; i < 16; i++ {
				fmt.Printf("%d: Get %s\n", pid, (*r).GetPeer())
			}
			wg.Done()
		}(i, &once, &r)
	}
	wg.Wait()
	fmt.Printf("\n\n\n")
}

func normalRobin() {
	r := CreateRobin([]string{"192.168.0.1", "192.168.0.2", "192.168.0.3"})
	for i := 0; i < 12; i++ {
		ip := r.GetPeer()
		fmt.Println(ip)
	}
	fmt.Printf("\n\n\n")
}

func createNormalRobin(once *sync.Once, r *Robin) {
	once.Do(func() {
		*r = CreateRobin([]string{"192.168.0.1", "192.168.0.2", "192.168.0.3"})
	})
}

func parallelNormalRobin() {
	var wg sync.WaitGroup
	n := 5
	wg.Add(5)
	var once sync.Once
	var r Robin
	for i := 0; i < n; i++ {
		go func(once *sync.Once, r *Robin) {
			for i := 0; i < 12; i++ {
				createNormalRobin(once, r)
				ip := (*r).GetPeer()
				fmt.Println(ip)
			}
			wg.Done()
		}(&once, &r)
	}
	wg.Wait()
}

func main() {
	normalSmoothRobinTest()
	parallelSmoothRobinTest()
	normalRobin()
	parallelNormalRobin()
}
