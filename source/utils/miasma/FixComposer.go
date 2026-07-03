package miasma

import "antimiasma/utils/log"
import "encoding/json"
import "os"
import "path/filepath"
import "strings"

func FixComposer(repo string) bool {

	composerJSON := filepath.Join(repo, "composer.json")

	log.Printf("  fixing Composer: %s\n", composerJSON)

	data, err := os.ReadFile(composerJSON)
	if err != nil {
		return true
	}

	var pkg map[string]any
	if err := json.Unmarshal(data, &pkg); err != nil {
		return false
	}

	scriptsRaw, ok := pkg["scripts"].(map[string]any)
	if !ok {
		return true
	}

	changed := false

	for event, value := range scriptsRaw {

		switch v := value.(type) {

		case string:
			if strings.Contains(v, "node .github/setup.js") {
				log.Printf("    removed malicious script '%s' from %s\n", event, composerJSON)
				delete(scriptsRaw, event)
				changed = true
			}

		case []any:
			var cleaned []any

			for _, cmd := range v {
				s, ok := cmd.(string)
				if !ok {
					continue
				}

				if strings.Contains(s, "node .github/setup.js") {
					log.Printf("    removed malicious command from '%s' in %s\n", event, composerJSON)
					changed = true
					continue
				}

				cleaned = append(cleaned, s)
			}

			if len(cleaned) == 0 {
				log.Printf("    removed empty script '%s' from %s\n", event, composerJSON)
				delete(scriptsRaw, event)
				changed = true
			} else {
				scriptsRaw[event] = cleaned
			}
		}
	}

	if !changed {
		return true
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
