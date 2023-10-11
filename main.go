package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	setupApi()

	log.Fatal(http.ListenAndServe(":3001", nil))
	// need to generate server.cert and server.key from openssl
	//log.Fatal(http.ListenAndServeTLS(":3001", "server.cert", "server.key"))
}

func setupApi() {

	ctx := context.Background()

	manager := NewManager(ctx)

	// create a hello world handler
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running"))
	}))
	http.HandleFunc("/ws", manager.serveWs)
	http.HandleFunc("/login", manager.loginHandler)
}
