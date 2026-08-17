package miasma

import utils_json "antimiasma/utils/json"
import "os"
import "path/filepath"

func FixGemini(repo string) bool {

	settings_path := filepath.Join(repo, ".gemini", "settings.json")

	data, err := os.ReadFile(settings_path)
	if err != nil {
		return true // nothing to fix
	}

	var settings map[string]any
	if err := utils_json.Unmarshal(data, &settings); err != nil {
		return false
	}

	hooks, ok := settings["hooks"]
	if !ok {
		return true // nothing to fix
	}

	cleaned, changed := utils_json.FilterKeyVal(hooks, "command", []string{
		"node .github/setup.js",
		"node .claude/setup.mjs",
		"node .gemini/setup.mjs",
		"node .vscode/setup.mjs",
	})
	if !changed {
		return true // already clean
	}

	settings["hooks"] = cleaned

	output, err := utils_json.Marshal(settings)
	if err != nil {
		return false
	}

	if err := os.WriteFile(settings_path, output, 0644); err != nil {
		return false
	}

	return true

}

