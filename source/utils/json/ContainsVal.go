package json

func ContainsVal(raw any, values []string) bool {

	switch data := raw.(type) {

	case string:

		for _, search := range values {

			if search == data {
				return true
			}

		}

	case map[string]any:

		contains := false

		for _, child := range data {

			child_contains := ContainsVal(child, values)

			if child_contains {
				contains = true
				break
			}

		}

		return contains

	case []any:

		contains := false

		for _, child := range data {

			child_contains := ContainsVal(child, values)

			if child_contains {
				contains = true
				break
			}

		}

		return contains

	}

	return false

}
