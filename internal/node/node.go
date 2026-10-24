package node

import "distributed-kv-storage/internal/storage"

type Node struct {
	ID    string
	Store *storage.Store
	

}

func NewNode(id string, store *storage.Store) *Node {
	return &Node{
		ID: id,
		Store: store,
	}
}

func (n *Node) Put(key, value string){
	n.Store.Put(key, value)
}

func (n *Node) Get(key string) (string, bool) {
	return n.Store.Get(key)
}

func (n *Node) Delete(key string){
	n.Store.Delete(key)
}