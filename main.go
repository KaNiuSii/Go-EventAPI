package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleFuncs()

	port := ":8081"
	fmt.Printf(" * Server running at: http://localhost%s * \n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleFuncs() {
	http.HandleFunc("/christmas/", CORS(christmas))
	http.HandleFunc("/halloween/", CORS(halloween))
	http.HandleFunc("/easter/", CORS(easter))
}
