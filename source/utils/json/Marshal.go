package json

import "bytes"
import "encoding/json"

func Marshal(raw any) ([]byte, error) {

	var buffer bytes.Buffer

	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")

	err0 := encoder.Encode(raw)

	if err0 == nil {
		return bytes.TrimSuffix(buffer.Bytes(), []byte{'\n'}), nil
	} else {
		return nil, err0
	}

}

