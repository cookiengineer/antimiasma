package miasma

func containsTasks(tasks []any, targets []string) bool {

	for _, t := range tasks {

		task, ok := t.(map[string]any)
		if !ok {
			continue
		}

		cmd, ok := task["command"].(string)
		if !ok {
			continue
		}

		for _, target := range targets {
			if cmd == target {
				return true
			}
		}

	}

	return false

}
