package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", WebhookServer)
	http.ListenAndServe(":8080", nil)
}

// WebhookServer responds to webhooks.
func WebhookServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ack %s", r.URL.Path[1:])
}
