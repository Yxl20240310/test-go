package main

import (
	"fmt"
	"net/http"
)

const PORT = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	fmt.Println("Listening on port", PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		.Println("Failed to start the server:", err)
	}
}
