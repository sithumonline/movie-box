package handler

import (
	"encoding/json"
	"net/http"
)

// RespondWithError return error message
func RespondWithError(w http.ResponseWriter, code int, msg interface{}) {
	RespondWithJSON(w, code, map[string]interface{}{"message": msg})
}

// RespondWithJSON write json response format
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	t := "data"

	if code > 300 {
		t = "errors"
	}

	response, _ := json.Marshal(map[string]interface{}{
		t: payload,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithText(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(msg))
}

func toInt(m json.Number) int {
	var n int
	json.Unmarshal([]byte(m), &n)
	return n
}
