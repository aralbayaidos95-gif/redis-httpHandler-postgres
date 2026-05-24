package main

import (
	"log"
	"net/http"
	"study/internal/http_handler"
	"study/internal/service"
	"study/internal/storage"
)

func main() {

	connStr := "postgres://postgres:2055@localhost:5432/postgres"

	store, err := storage.NewStorage(connStr)

	if err != nil {
		log.Fatal(err)
	}
	connStr = "127.0.0.1:6379"

	redis := storage.NewRedis(connStr)

	svc := service.NewService(store, redis)

	handler := http_handler.NewHandler(svc)

	http.HandleFunc("/user/post", handler.PostUsers)
	http.HandleFunc("/user/get", handler.GetUser)
	http.HandleFunc("/user/get/all", handler.GetUsers)
	http.HandleFunc("/user/delete", handler.DeleteUser)
	http.HandleFunc("/user/put", handler.PutUser)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("error http listen and serve", err)
	}
}
