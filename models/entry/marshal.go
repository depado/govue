package entry

import "encoding/json"

// Encode dumps an Entry to json.
func (e Entry) Encode() ([]byte, error) {
	enc, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

// Decode loads an Entry from json
func (e *Entry) Decode(data []byte) error {
	return json.Unmarshal(data, e)
}
