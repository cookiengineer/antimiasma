package json

func ContainsKeyVal(raw any, key string, values []string) bool {

	switch data := raw.(type) {
	
	case map[string]any:

		if value, ok := data[key].(string); ok {

			for _, search := range values {

				if search == value {
					return true
				}

			}

		}

		contains := false

		for _, child := range data {

			child_contains := ContainsKeyVal(child, key, values)

			if child_contains {
				contains = true
				break
			}

		}

		return contains

	case []any:

		contains := false

		for _, child := range data {

			child_contains := ContainsKeyVal(child, key, values)

			if child_contains {
				contains = true
				break
			}

		}

		return contains

	}

	return false

}
