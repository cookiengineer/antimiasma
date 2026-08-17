package js

import "strings"

func IsSuspicious(data []byte) bool {

	result  := false
	content := strings.TrimSpace(string(data))

	if strings.Contains(content, "eval") && strings.Contains(content, "String.fromCharCode") {
		result = true
	}

	return result

}
