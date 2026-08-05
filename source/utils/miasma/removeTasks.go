package miasma

func removeTasks(tasks []any, targets []string) (any, bool) {

	changed := false
	result := make([]any, 0, len(tasks))

	for _, t := range tasks {

		task, ok := t.(map[string]any)
		if !ok {
			result = append(result, t)
			continue
		}

		cmd, ok := task["command"].(string)
		if ok {

			remove := false

			for _, target := range targets {
				if cmd == target {
					remove = true
					break
				}
			}

			if remove {
				changed = true
				continue
			}

		}

		result = append(result, task)

	}

	return result, changed

}
