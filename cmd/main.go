package main

import (
	"log"
	"net/http"
	"study/internal/http_handler"
	"study/internal/service"
	"study/internal/storage"
)

func main() {

	connStr := "postgres://postgres:2055@localhost:5432/mydb"

	store, err := storage.NewStorage(connStr)

	if err != nil {
		log.Fatal(err)
	}

	svc := service.NewService(store)

	handler := http_handler.NewHandler(svc)

	http.HandleFunc("/user/post", handler.PostUsers)
	http.HandleFunc("/user/get", handler.GetUser)
	http.HandleFunc("/user/delete", handler.DeleteUser)
	http.HandleFunc("/user/put", handler.PutUser)

	http.ListenAndServe("8080", nil)
}
