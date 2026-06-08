package miasma

import "encoding/json"
import "os"
import "path/filepath"
import "strings"

func FixComposer(repo string) bool {

	composerJSON := filepath.Join(repo, "composer.json")

	data, err := os.ReadFile(composerJSON)
	if err != nil {
		return true // nothing to fix
	}

	var pkg map[string]any
	if err := json.Unmarshal(data, &pkg); err != nil {
		return false
	}

	scriptsRaw, ok := pkg["scripts"].(map[string]any)
	if !ok {
		return true // nothing to fix
	}

	changed := false

	for event, value := range scriptsRaw {

		switch v := value.(type) {

		case string:
			if strings.Contains(v, "node .github/setup.js") {
				delete(scriptsRaw, event)
				changed = true
			}

		case []any:
			// Composer allows multiple commands per event
			var cleaned []any

			for _, cmd := range v {
				s, ok := cmd.(string)
				if !ok {
					continue
				}

				if strings.Contains(s, "node .github/setup.js") {
					changed = true
					continue
				}

				cleaned = append(cleaned, s)
			}

			if len(cleaned) == 0 {
				delete(scriptsRaw, event)
				changed = true
			} else {
				scriptsRaw[event] = cleaned
			}
		}
	}

	if !changed {
		return true // already clean
	}

	output, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return false
	}

	if err := os.WriteFile(composerJSON, output, 0644); err != nil {
		return false
	}

	return true
}
