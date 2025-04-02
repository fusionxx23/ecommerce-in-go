package main

import (
	"fmt"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/database"
)

func init() {
	database.ConnectDatabase()
	database.SyncDb()
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from the API!")
	})

	http.ListenAndServe(":8080", nil)
}
