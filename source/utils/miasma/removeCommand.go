package miasma

func removeCommand(v any, target string) (any, bool) {

	switch x := v.(type) {

	case map[string]any:
		changed := false

		if cmd, ok := x["command"].(string); ok && cmd == target {
			return nil, true
		}

		result := make(map[string]any, len(x))

		for k, value := range x {
			cleaned, childChanged := removeCommand(value, target)

			if childChanged {
				changed = true
			}

			if cleaned != nil {
				result[k] = cleaned
			}
		}

		return result, changed

	case []any:
		changed := false
		result := make([]any, 0, len(x))

		for _, value := range x {
			cleaned, childChanged := removeCommand(value, target)

			if childChanged {
				changed = true
			}

			if cleaned != nil {
				result = append(result, cleaned)
			}
		}

		return result, changed
	}

	return v, false

}
