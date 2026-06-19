package pkgbuild

import "fmt"
import "strings"

func RemoveDepends(content string, search string) string {

	index := -1

	if tmp := strings.Index(content, "depends="); tmp >= 0 {
		index = tmp + len("depends=")
	} else if tmp := strings.Index(content, "depends ="); tmp >= 0 {
		index = tmp + len("depends =")
	}

	if index < 0 {
		return content
	}

	for index < len(content) && (content[index] == ' ' || content[index] == '\t') {
		index++
	}

	if index >= len(content) || content[index] != '(' {
		return content
	}

	end := findMatching(content[index:], '(', ')')

	if end < 0 {
		return content
	}

	section := content[index : index+end+1]

	quoted_single := fmt.Sprintf("'%s'", search)
	quoted_double := fmt.Sprintf("\"%s\"", search)

	match_index := -1
	match_len   := 0

	if tmp := strings.Index(section, quoted_single); tmp >= 0 {
		match_index = tmp
		match_len   = len(quoted_single)
	} else if tmp := strings.Index(section, quoted_double); tmp >= 0 {
		match_index = tmp
		match_len   = len(quoted_double)
	}

	if match_index < 0 {
		return content
	}

	abs_start := index + match_index
	abs_end   := index + match_index + match_len

	result := content[:abs_start] + content[abs_end:]

	result = strings.ReplaceAll(result, "  ", " ")
	result = strings.ReplaceAll(result, "( ", "(")
	result = strings.ReplaceAll(result, " )", ")")

	return result

}
