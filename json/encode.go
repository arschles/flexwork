package json

import (
	"encoding/json"
	"net/http"

	"github.com/arschles/flexwork"
)

// Encode returns a response that encodes i to JSON
func Encode(code int, i interface{}) flexwork.Response {
	return flexwork.ResponseFunc(func(w http.ResponseWriter) {
		b, err := json.Marshal(i)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(jsonEncodeErr))
			return
		}
		w.WriteHeader(code)
		w.Write(b)
	})
}
