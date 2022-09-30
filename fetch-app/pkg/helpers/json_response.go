package helpers

import (
	"encoding/json"
	"net/http"
)

func BuildJSON(rw http.ResponseWriter, data interface{}) ([]byte, error) {
	rw.Header().Add("Content-Type", "application/json")
	response, err := json.Marshal(struct {
		Data interface{} `json:"data"`
	}{
		Data: data,
	})
	if err != nil {
		return []byte{}, err
	}

	return response, nil
}
