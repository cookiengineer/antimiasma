package pkgbuild

import "fmt"
import "strings"

func HasDepends(content string, search string) bool {

	index := -1

	if tmp := strings.Index(content, "depends="); tmp >= 0 {
		index = tmp + len("depends=")
	} else if tmp := strings.Index(content, "depends ="); tmp >= 0 {
		index = tmp + len("depends =")
	}

	if index < 0 {
		return false
	}

	for index < len(content) && (content[index] == ' ' || content[index] == '\t') {
		index++
	}

	if index >= len(content) || content[index] != '(' {
		return false
	}

	end := findMatching(content[index:], '(', ')')

	if end < 0 {
		return false
	}

	section := content[index : index+end+1]

	if strings.Contains(section, fmt.Sprintf("'%s'", search)) {
		return true
	}

	if strings.Contains(section, fmt.Sprintf("\"%s\"", search)) {
		return true
	}

	return false

}
