package gome

import (
	"encoding/json"
	"os"
)

// ReadJSON reads a JSON file at the specified path into an interface.
func ReadJSON(src string, data interface{}) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	return json.NewDecoder(f).Decode(data)
}

// SaveJSON encodes an interface into JSON and saves it to a file.
func WriteJSON(dst string, data interface{}) error {
	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(data)
}
