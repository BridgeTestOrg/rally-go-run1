// Command rally serves the Rally website.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BridgeTestOrg/rally-starter-go/internal/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, web.NewRouter()))
}
