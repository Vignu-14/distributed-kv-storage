package transport

import (
	"encoding/json"
	"net/http"
)

func(s *Server) handleHealth(w http.ResponseWriter, r *http.Request){
	writeJson(w, http.StatusOK, KVResponse{
		NodeId: s.Node.ID,
		Status: "alive",
	})
}
  

func writeJson(w http.ResponseWriter, code int, payload any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}