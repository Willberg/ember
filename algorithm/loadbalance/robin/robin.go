package robin

import (
	"sync"
)

type Robin interface {
	GetPeer() string
}

type Pair struct {
	Ip     string
	Weight int
}

type peer struct {
	ip                      string
	weight, effectiveWeight int
}

type smoothRobin struct {
	peers []peer
	pos   int
	total int
	mu    sync.Mutex
}

func CreateSmoothRobin(pairs []Pair) Robin {
	var ps []peer
	pos, total := 0, 0
	for i, p := range pairs {
		if p.Weight > pairs[i].Weight {
			pos = i
		}
		total += p.Weight
		ps = append(ps, peer{p.Ip, p.Weight, p.Weight})
	}
	return &smoothRobin{peers: ps, pos: pos, total: total}
}

// 平滑轮询
func (r *smoothRobin) GetPeer() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	ip := r.peers[r.pos].ip
	r.peers[r.pos].effectiveWeight -= r.total
	pos := 0
	for i, p := range r.peers {
		curWeight := p.effectiveWeight + p.weight
		r.peers[i].effectiveWeight = curWeight
		if curWeight > r.peers[pos].effectiveWeight {
			pos = i
		}
	}
	r.pos = pos
	return ip
}

type robin struct {
	ips []string
	pos int
	mu  sync.Mutex
}

// 简单轮询
func (r *robin) GetPeer() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	n := len(r.ips)
	pos := (r.pos + 1) % n
	r.pos = pos
	return r.ips[pos]
}

func CreateRobin(ips []string) Robin {
	return &robin{ips: ips, pos: -1}
}
