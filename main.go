package main

import (
	"io"
	"net/http"
	"os"
	"time"
	"fmt"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":"+port, nil)
}