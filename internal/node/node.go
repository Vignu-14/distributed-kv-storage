package node

import "distributed-kv-storage/internal/storage"

type Node struct {
	ID    string
	store *storage.Store

}

func NewNode(id string, store *storage.Store) *Node {
	return &Node{
		ID: id,
		store: store,
	}
}

func (n *Node) Put()