package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, error := fmt.Fprintf(w, "Hello World")
		if error != nil {
			fmt.Println(error)
		}
		fmt.Printf("Wrote %d bytes to the connection\n", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
