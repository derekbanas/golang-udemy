package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,
		r *http.Request) {
		nb, err := fmt.Fprintf(w, "Hello Browser")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes Written : %d", nb)
	})
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
