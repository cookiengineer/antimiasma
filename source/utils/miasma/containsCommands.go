package miasma

func containsCommands(v any, targets []string) bool {

	switch x := v.(type) {

	case map[string]any:

		if cmd, ok := x["command"].(string); ok {
			for _, target := range targets {
				if cmd == target {
					return true
				}
			}
		}

		for _, value := range x {
			if containsCommands(value, targets) {
				return true
			}
		}

	case []any:

		for _, value := range x {
			if containsCommands(value, targets) {
				return true
			}
		}

	}

	return false

}
