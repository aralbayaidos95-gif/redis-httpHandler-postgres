package http_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"study/internal/models"
)

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.svc.DeleteUser(context.Background(), user.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("deleted")
}
