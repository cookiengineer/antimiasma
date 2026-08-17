package miasma

import utils_json "antimiasma/utils/json"
import "os"
import "path/filepath"

func FixVSCode(repo string) bool {

	tasks_path := filepath.Join(repo, ".vscode", "tasks.json")

	data, err := os.ReadFile(tasks_path)
	if err != nil {
		return true // nothing to fix
	}

	var cfg map[string]any
	if err := utils_json.Unmarshal(data, &cfg); err != nil {
		return false
	}

	tasks, ok := cfg["tasks"].([]any)
	if !ok {
		return true // nothing to fix
	}

	cleaned, changed := utils_json.FilterKeyVal(tasks, "command", []string{
		"node .github/setup.js",
		"node .claude/setup.mjs",
		"node .gemini/setup.mjs",
		"node .vscode/setup.mjs",
	})
	if !changed {
		return true // already clean
	}

	cfg["tasks"] = cleaned

	output, err := utils_json.Marshal(cfg)
	if err != nil {
		return false
	}

	if err := os.WriteFile(tasks_path, output, 0644); err != nil {
		return false
	}

	return true
}
