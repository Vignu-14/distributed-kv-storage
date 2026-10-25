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
  
func(s *Server) handleKV(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodPost:
		var req PutRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid JSON body: key and value required", http.StatusBadRequest)
			return
		}
		s.Node.Store.Put(req.Key, req.Value)
		writeJson(w, http.StatusCreated, KVResponse{
			NodeId: s.Node.ID,
			Key: req.Key,
			Value: req.Value,
			Status: "created",
		})
	
	}
}

func writeJson(w http.ResponseWriter, code int, payload any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}