package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	aport := 8080
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Listening... ", aport)

	err := http.ListenAndServe(":"+strconv.Itoa(aport), nil)
	if err != nil {
		log.Fatal("error", err)
	}
}
