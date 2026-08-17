package json

func FilterVal(raw any, values []string) (any, bool) {

	switch data := raw.(type) {

	case string:

		for _, search := range values {

			if search == data {
				return nil, true
			}

		}

	case map[string]any:

		changed := false
		result  := make(map[string]any, len(data))

		for c, child := range data {

			cleaned, child_changed := FilterVal(child, values)

			if child_changed {
				changed = true
			}

			if cleaned != nil {
				result[c] = cleaned
			}

		}

		return result, changed

	case []any:

		changed := false
		result  := make([]any, 0, len(data))

		for _, child := range data {

			cleaned, child_changed := FilterVal(child, values)

			if child_changed {
				changed = true
			}

			if cleaned != nil {
				result = append(result, cleaned)
			}

		}

		return result, changed

	}

	return raw, false

}
