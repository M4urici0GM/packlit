package main

import (
	"encoding/json"
	"log"
	"net/http"

	packlit "github.com/m4urici0gm/packlit/pkg"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req packlit.PackagerOptions
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sp, err := packlit.BuildFromJSON(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		command, err := sp.PreviewCommand()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"command_preview": command,
		})
	})

	log.Println("Starting server on :8080. POST /")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
