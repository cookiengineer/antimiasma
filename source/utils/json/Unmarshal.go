package json

import "encoding/json"

func Unmarshal(data []byte, value any) error {
	return json.Unmarshal(data, value)
}
