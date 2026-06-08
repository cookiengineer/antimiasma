package miasma

import "encoding/json"
import "os"
import "path/filepath"

func FixVSCode(repo string) bool {

	tasksPath := filepath.Join(repo, ".vscode", "tasks.json")

	data, err := os.ReadFile(tasksPath)
	if err != nil {
		return true // nothing to fix
	}

	var cfg map[string]any
	if err := json.Unmarshal(data, &cfg); err != nil {
		return false
	}

	rawTasks, ok := cfg["tasks"].([]any)
	if !ok {
		return true // nothing to fix
	}

	changed := false
	cleanedTasks := make([]any, 0, len(rawTasks))

	for _, t := range rawTasks {
		task, ok := t.(map[string]any)
		if !ok {
			continue
		}

		cmd, _ := task["command"].(string)

		if cmd == "node .github/setup.js" {
			changed = true
			continue // remove this task
		}

		cleanedTasks = append(cleanedTasks, task)
	}

	if !changed {
		return true // already clean
	}

	cfg["tasks"] = cleanedTasks

	output, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return false
	}

	if err := os.WriteFile(tasksPath, output, 0644); err != nil {
		return false
	}

	return true
}
