package miasma

import "encoding/json"
import "os"
import "path/filepath"

func FixGemini(repo string) bool {

	settingsPath := filepath.Join(repo, ".gemini", "settings.json")

	data, err := os.ReadFile(settingsPath)
	if err != nil {
		return true // nothing to fix
	}

	var settings map[string]any
	if err := json.Unmarshal(data, &settings); err != nil {
		return false
	}

	hooks, ok := settings["hooks"]
	if !ok {
		return true // nothing to fix
	}

	cleaned, changed := removeCommands(hooks, []string{
		"node .github/setup.js",
		"node .claude/setup.mjs",
		"node .gemini/setup.mjs",
		"node .vscode/setup.mjs",
	})
	if !changed {
		return true // already clean
	}

	settings["hooks"] = cleaned

	output, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return false
	}

	if err := os.WriteFile(settingsPath, output, 0644); err != nil {
		return false
	}

	return true

}

