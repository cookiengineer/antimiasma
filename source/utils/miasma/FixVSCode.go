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

	tasks, ok := cfg["tasks"].([]any)
	if !ok {
		return true // nothing to fix
	}

	cleaned, changed := removeTasks(tasks, []string{
		"node .github/setup.js",
		"node .claude/setup.mjs",
		"node .gemini/setup.mjs",
		"node .vscode/setup.mjs",
	})
	if !changed {
		return true // already clean
	}

	cfg["tasks"] = cleaned

	output, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return false
	}

	if err := os.WriteFile(tasksPath, output, 0644); err != nil {
		return false
	}

	return true
}
