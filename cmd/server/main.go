package main

import (
	"log"
	"net/http"

	"github.com/glhou/logm/internal/web"
)

func main() {
	mux := http.NewServeMux()

	web.RegisterRoutes(mux)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
