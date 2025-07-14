package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"
)

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	Result int `json:"result"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		base64EncodedPayload := r.URL.Query().Get("payload")

		payload, err := base64.StdEncoding.DecodeString(base64EncodedPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req Request
		if err := json.Unmarshal(payload, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		time.Sleep(5 * time.Second)

		result := req.A * req.B

		json.NewEncoder(w).Encode(&Response{Result: result})
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}
