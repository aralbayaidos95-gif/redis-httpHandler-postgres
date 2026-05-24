package http_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"study/internal/models"
)

func (h *Handler) PostUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	if err := h.svc.CreatUser(context.Background(), user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("created")
}
