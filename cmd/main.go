package main

import (
	"log"
	"net/http"
	"os"
	"study/internal/http_handler"
	"study/internal/service"
	"study/internal/storage"
)

func main() {

	CONN_STR := os.Getenv("CONN_STR")
	CONN_STR_RDB := os.Getenv("CONN_STR_RDB")

	store, err := storage.NewStorage(CONN_STR)
	redis := storage.NewRedis(CONN_STR_RDB)

	if err != nil {
		log.Fatal(err)
	}

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
