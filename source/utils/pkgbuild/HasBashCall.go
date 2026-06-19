package pkgbuild

import "strings"

func HasBashCall(content string, function string, args string) bool {

	index := findFunctionStart(content, function)

	if index < 0 {
		return false
	}

	rest  := content[index:]
	brace := strings.Index(rest, "{")

	if brace < 0 {
		return false
	}

	end := findMatching(rest[brace:], '{', '}')

	if end < 0 {
		return false
	}

	body := rest[brace : brace+end+1]

	if strings.Contains(body, args) {
		return true
	}

	return false

}
