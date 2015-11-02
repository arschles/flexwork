package json

import (
	"encoding/json"
	"io"
)

// Decode is a JSON decoding convenience func. Sample usage:
//
//  if err := json.Decode(request.Body, &thing); err != nil {
//    return json.Error(http.StatusBadRequest, json.ErrFormatting, "invalid json")
//  }
func Decode(rc io.ReadCloser, into interface{}) error {
	if err := json.NewDecoder(rc).Decode(into); err != nil {
		return err
	}
	rc.Close()
	return nil
}
