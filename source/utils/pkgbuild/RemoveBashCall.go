package pkgbuild

import "strings"

func RemoveBashCall(content string, function string, args string) string {

	index := findFunctionStart(content, function)

	if index < 0 {
		return content
	}

	rest  := content[index:]
	brace := strings.Index(rest, "{")

	if brace < 0 {
		return content
	}

	end := findMatching(rest[brace:], '{', '}')

	if end < 0 {
		return content
	}

	body := rest[brace : brace+end+1]

	if len(body) < 2 {
		return content
	}

	inner := body[1 : len(body)-1]
	lines := strings.Split(inner, "\n")

	removed  := false
	filtered := make([]string, 0, len(lines))

	for l := 0; l < len(lines); l++ {
		line    := lines[l]
		trimmed := strings.TrimLeft(line, " \t")

		if strings.HasPrefix(trimmed, args) {
			removed = true
		} else {
			filtered = append(filtered, line)
		}
	}

	if !removed {
		return content
	}

	new_inner := strings.Join(filtered, "\n")
	new_body  := "{" + new_inner + "}"

	result := content[:index] + rest[:brace] + new_body + content[index+brace+end+1:]

	return result

}
