package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/setting", func(w http.ResponseWriter, r *http.Request) {
		setting := struct {
			Frequency int
		}{
			Frequency: 0,
		}
		json.NewEncoder(w).Encode(setting)
	})
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
