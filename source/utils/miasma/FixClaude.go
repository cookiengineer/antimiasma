package miasma

import "antimiasma/utils/log"
import "encoding/json"
import "os"
import "path/filepath"

func FixClaude(repo string) bool {

	settingsPath := filepath.Join(repo, ".claude", "settings.json")

	log.Printf("  fixing Claude: %s\n", settingsPath)

	data, err := os.ReadFile(settingsPath)
	if err != nil {
		return true
	}

	var settings map[string]any
	if err := json.Unmarshal(data, &settings); err != nil {
		return false
	}

	hooks, ok := settings["hooks"]
	if !ok {
		return true
	}

	cleaned, changed := removeCommand(hooks, "node .github/setup.js")
	if !changed {
		return true
	}

	log.Printf("    cleaned hooks in %s\n", settingsPath)

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

