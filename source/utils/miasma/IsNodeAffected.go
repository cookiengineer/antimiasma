package miasma

import "encoding/json"
import "os"
import "path/filepath"
import "strings"

func IsNodeAffected(repo string) bool {

	packageJSON := filepath.Join(repo, "package.json")

	data, err := os.ReadFile(packageJSON)
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
		script, ok := value.(string)
		if !ok {
			continue
		}

		if strings.Contains(script, "node .github/setup.js") {
			return true
		}

		if strings.Contains(script, "src/hooks/deps") {
			return true
		}

	}

	return false

}
