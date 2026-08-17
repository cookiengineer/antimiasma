package miasma

import utils_json "antimiasma/utils/json"
import "os"
import "path/filepath"

func IsVSCodeAffected(repo string) bool {

	tasksPath := filepath.Join(repo, ".vscode", "tasks.json")

	data, err := os.ReadFile(tasksPath)
	if err != nil {
		return false
	}

	var cfg map[string]any
	if err := utils_json.Unmarshal(data, &cfg); err != nil {
		return false
	}

	tasks, ok := cfg["tasks"].([]any)
	if !ok {
		return false
	}

	return utils_json.ContainsKeyVal(tasks, "command", []string{
		"node .github/setup.js",
		"node .claude/setup.mjs",
		"node .gemini/setup.mjs",
		"node .vscode/setup.mjs",
	})

}
