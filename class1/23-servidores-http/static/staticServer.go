package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

    log.Println("Executando...")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
