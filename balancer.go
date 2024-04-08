package main

import "sync"

type Balancer struct {
	index   int
	mu      sync.Mutex
	servers []Server
}

func NewBalancer(servers []Server) *Balancer {
	return &Balancer{
		index:   0,
		servers: servers,
	}
}

func (b *Balancer) getServer() Server {
	defer func() {
		b.mu.Unlock()
	}()

	b.mu.Lock()
	server := b.servers[b.index]
	b.index = (b.index + 1) % len(b.servers)

	return server
}
