package main

import (
        "fmt"
        "net/http"

	"github.com/mattes/go-asciibot"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, asciibot.Random())
}
