package main

import (
	"github.com/SamWheating/battlesnake2020/pkg/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Snake Server started on port %s", port)

	http.HandleFunc("/start", handlers.Start)
	http.HandleFunc("/end", handlers.End)
	http.HandleFunc("/move", handlers.Move)
	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/", handlers.Index)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
