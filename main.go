package main

import (
	"net/http"
)

func main() {
	sm := http.NewServeMux()

	sm.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Project is working"))
	})

	http.ListenAndServe(":8080", sm)

}
