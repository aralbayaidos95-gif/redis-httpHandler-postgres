package http_handler

import (
	"context"
	"encoding/json"
	"net/http"
)

func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.GetUsers(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
