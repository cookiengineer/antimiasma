package miasma

import "encoding/json"
import "os"
import "path/filepath"

func IsGeminiAffected(repo string) bool {

	settingsPath := filepath.Join(repo, ".gemini", "settings.json")

	data, err := os.ReadFile(settingsPath)
	if err != nil {
		return false
	}

	var settings map[string]any
	if err := json.Unmarshal(data, &settings); err != nil {
		return false
	}

	hooks, ok := settings["hooks"]
	if !ok {
		return false
	}

	return containsCommands(hooks, []string{
		"node .github/setup.js",
		"node .claude/setup.mjs",
		"node .gemini/setup.mjs",
		"node .vscode/setup.mjs",
	})

}

