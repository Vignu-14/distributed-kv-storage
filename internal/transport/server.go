package transport

import (
	"distributed-kv-storage/internal/node"
	"net/http"
)

type Server struct {
	Node *node.Node
	mux  *http.ServeMux
}

func NewServer(node *node.Node) *Server {
	s := &Server{
		Node: node,
		mux: http.NewServeMux(),
	}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes(){
	s.mux.HandleFunc("/kv", s.handleKV)
	s.mux.HandleFunc("/internal/health", s.handleHealth)
}
