package miasma

import utils_json "antimiasma/utils/json"
import "os"
import "path/filepath"

func IsGeminiAffected(repo string) bool {

	settings_path := filepath.Join(repo, ".gemini", "settings.json")

	data, err := os.ReadFile(settings_path)
	if err != nil {
		return false
	}

	var settings map[string]any
	if err := utils_json.Unmarshal(data, &settings); err != nil {
		return false
	}

	hooks, ok := settings["hooks"]
	if !ok {
		return false
	}

	return utils_json.ContainsKeyVal(hooks, "command", []string{
		"node .github/setup.js",
		"node .claude/setup.mjs",
		"node .gemini/setup.mjs",
		"node .vscode/setup.mjs",
	})

}

