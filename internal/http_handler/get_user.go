package http_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"study/internal/models"
)

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	user, err := h.svc.GetUser(context.Background(), user.Name)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)

}
