package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Disable compression until we figure out what's wrong.
	//http.HandleFunc("/", compress(moomoo))
	http.HandleFunc("/", moomoo)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
