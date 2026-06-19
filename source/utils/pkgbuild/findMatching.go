package pkgbuild

func findMatching(content string, open_token byte, close_token byte) int {

	depth := 0

	for c := 0; c < len(content); c++ {

		if content[c] == open_token {

			depth++

		} else if content[c] == close_token {

			depth--

			if depth == 0 {
				return c
			}

		}

	}

	return -1

}

