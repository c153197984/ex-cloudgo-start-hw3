package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// This server send "Hello" to client based on URL.
	//
	// localhost:8080/test/use
	// > Hello test use!
	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s!\n",
			strings.Join(strings.Split(r.URL.String(), "/")[1:], " "))
	})

	fmt.Println("Listening to port 8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
