package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	req int32 = 0
)

type LoadBalancer struct {
	client []*Client
	size   int32
}

type Client struct {
	name string
}

func NewBalancer(size int32) *LoadBalancer {
	loadBalancer := &LoadBalancer{client: make([]*Client, size), size: size}
	for i := 0; i < int(size); i++ {
		loadBalancer.client[i] = &Client{name: fmt.Sprintf("client:%d", i)}
	}
	return loadBalancer
}

// 随机算法
func (l *LoadBalancer) getClientRand() *Client {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Int31n(100) % l.size
	return l.client[idx]
}

// 轮询算法
func (l *LoadBalancer) getClientRoundRobin() *Client {
	req := atomic.AddInt32(&req, 1)
	return l.client[req%l.size]
}

func (c *Client) Do() {
	fmt.Println(c.name)
}

func main() {
	loadBalance := NewBalancer(4)

	fmt.Println("随机:")
	for i := 0; i < 10; i++ {
		client := loadBalance.getClientRand()
		client.Do()
	}

	fmt.Println("轮询:")
	var n sync.WaitGroup
	for i := 0; i < 10; i++ {
		n.Add(1)
		go func() {
			client := loadBalance.getClientRoundRobin()
			client.Do()
			n.Done()
		}()
		n.Wait()
	}
}
