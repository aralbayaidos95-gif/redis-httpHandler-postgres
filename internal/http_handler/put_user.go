package http_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"study/internal/models"
)

func (h Handler) PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	if err := h.svc.UpdateUser(context.Background(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("updated")
}
