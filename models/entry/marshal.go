package entry

import "encoding/json"

// Encode dumps an Entry to json.
func (e Entry) Encode() ([]byte, error) {
	return json.Marshal(e)
}

// Decode loads an Entry from json
func (e *Entry) Decode(data []byte) error {
	return json.Unmarshal(data, e)
}
