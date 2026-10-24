package cluster

import (
	"sync"
	"time"
)

type Peer struct{
	ID string
	Addr string
	alive bool
	lastSeen time.Time
	mu sync.RWMutex
}

func NewPeer(id, addr string) *Peer {
	return &Peer{
		ID: id,
		Addr: addr,
		alive: false,
	}
}

func (p *Peer) SetAlive(status bool)bool{
	p.mu.Lock()
	changed := p.alive != status
	if status {
		p.lastSeen = time.Now()
	}
	p.mu.Unlock()
	return changed
}

func (p *Peer)IsAlive() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return  p.alive
}