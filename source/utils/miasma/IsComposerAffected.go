package miasma

import "antimiasma/utils/log"
import "encoding/json"
import "os"
import "path/filepath"
import "strings"

func IsComposerAffected(repo string) bool {

	composerJSON := filepath.Join(repo, "composer.json")

	log.Printf("  checking Composer: %s\n", composerJSON)

	data, err := os.ReadFile(composerJSON)
	if err != nil {
		return false
	}

	var pkg map[string]any
	if err := json.Unmarshal(data, &pkg); err != nil {
		return false
	}

	scripts, ok := pkg["scripts"].(map[string]any)
	if !ok {
		return false
	}

	for _, value := range scripts {

		switch v := value.(type) {

		case string:
			if strings.Contains(v, "node .github/setup.js") {
				return true
			}

		case []any:
			for _, cmd := range v {
				s, ok := cmd.(string)
				if !ok {
					continue
				}

				if strings.Contains(s, "node .github/setup.js") {
					return true
				}
			}
		}
	}

	return false
}
