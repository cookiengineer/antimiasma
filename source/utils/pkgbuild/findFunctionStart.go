package pkgbuild

import "strings"

func findFunctionStart(content string, name string) int {

	for i := 0; i < len(content)-len(name); i++ {

		if content[i:i+len(name)] == name {

			rest := content[i+len(name):]
			trimmed := strings.TrimLeft(rest, " \t")

			if strings.HasPrefix(trimmed, "()") {
				return i
			}

		}

	}

	return -1

}

