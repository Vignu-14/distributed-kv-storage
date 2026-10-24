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

