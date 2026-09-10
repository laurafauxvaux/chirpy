package main

import "net/http"

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.Handle("/", http.FileServer(http.Dir(".")))

	server.ListenAndServe()
}
