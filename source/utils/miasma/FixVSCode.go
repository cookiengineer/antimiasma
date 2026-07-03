package miasma

import "antimiasma/utils/log"
import "encoding/json"
import "os"
import "path/filepath"

func FixVSCode(repo string) bool {

	tasksPath := filepath.Join(repo, ".vscode", "tasks.json")

	log.Printf("  fixing VSCode: %s\n", tasksPath)

	data, err := os.ReadFile(tasksPath)
	if err != nil {
		return true
	}

	var cfg map[string]any
	if err := json.Unmarshal(data, &cfg); err != nil {
		return false
	}

	rawTasks, ok := cfg["tasks"].([]any)
	if !ok {
		return true
	}

	changed := false
	cleanedTasks := make([]any, 0, len(rawTasks))

	for _, t := range rawTasks {
		task, ok := t.(map[string]any)
		if !ok {
			continue
		}

		cmd, _ := task["command"].(string)

		if cmd == "node .github/setup.js" {
			changed = true
			continue
		}

		cleanedTasks = append(cleanedTasks, task)
	}

	if !changed {
		return true
	}

	log.Printf("    cleaned tasks in %s\n", tasksPath)

	cfg["tasks"] = cleanedTasks

	output, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return false
	}

	if err := os.WriteFile(tasksPath, output, 0644); err != nil {
		return false
	}

	return true
}
