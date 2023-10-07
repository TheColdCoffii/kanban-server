package api

import (
	"encoding/json"
	"net/http"

	"github.com/E-V-Castillo/kanban-api/internal/types"
)

func WriteJson(w http.ResponseWriter, jsonResponse types.JSONResponse, code int) {
	res, err := json.Marshal(jsonResponse)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, writeErr := w.Write(res)
	if writeErr != nil {
		panic(err)
	}
}
