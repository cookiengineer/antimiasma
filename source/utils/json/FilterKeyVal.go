package json

func FilterKeyVal(raw any, key string, values []string) (any, bool) {

	switch data := raw.(type) {

	case map[string]any:

		if value, ok := data[key].(string); ok {

			for _, search := range values {

				if search == value {
					return nil, true
				}

			}

		}

		changed := false
		result  := make(map[string]any, len(data))

		for c, child := range data {

			cleaned, child_changed := FilterKeyVal(child, key, values)

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

			cleaned, child_changed := FilterKeyVal(child, key, values)

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
