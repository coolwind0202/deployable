package main

import (
	"log"
	"net/http"

	"github.com/coolwind0202/deployable/components"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		components.Hello("John").Render(req.Context(), w)
	})

	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
