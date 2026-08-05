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

	return containsTasks(tasks, []string{
		"node .github/setup.js",
		"node .claude/setup.mjs",
		"node .gemini/setup.mjs",
		"node .vscode/setup.mjs",
	})

}
