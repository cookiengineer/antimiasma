package miasma

func containsCommand(v any, target string) bool {

	switch x := v.(type) {
	case map[string]any:
		if cmd, ok := x["command"].(string); ok && cmd == target {
			return true
		}

		for _, value := range x {
			if containsCommand(value, target) {
				return true
			}
		}

	case []any:
		for _, value := range x {
			if containsCommand(value, target) {
				return true
			}
		}
	}

	return false

}
