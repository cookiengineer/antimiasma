package miasma

import "encoding/json"
import "os"
import "path/filepath"

func IsVSCodeAffected(repo string) bool {

	tasksPath := filepath.Join(repo, ".vscode", "tasks.json")

	data, err := os.ReadFile(tasksPath)
	if err != nil {
		return false
	}

	var cfg map[string]any
	if err := json.Unmarshal(data, &cfg); err != nil {
		return false
	}

	tasks, ok := cfg["tasks"].([]any)
	if !ok {
		return false
	}

	for _, t := range tasks {
		task, ok := t.(map[string]any)
		if !ok {
			continue
		}

		cmd, ok := task["command"].(string)
		if ok && cmd == "node .github/setup.js" {
			return true
		}
	}

	return false
}
