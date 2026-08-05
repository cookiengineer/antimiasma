package miasma

import "encoding/json"
import "os"
import "path/filepath"
import "strings"

func FixNode(repo string) bool {

	packageJSON := filepath.Join(repo, "package.json")

	data, err := os.ReadFile(packageJSON)
	if err != nil {
		return true // nothing to fix
	}

	var pkg map[string]any
	if err := json.Unmarshal(data, &pkg); err != nil {
		return false
	}

	scripts, ok := pkg["scripts"].(map[string]any)
	if !ok {
		return true // nothing to fix
	}

	changed := false

	for name, value := range scripts {

		script, ok := value.(string)

		if !ok {
			continue
		}

		if strings.Contains(script, "node .github/setup.js") {
			delete(scripts, name)
			changed = true
		} else if strings.Contains(script, "src/hooks/deps") {
			delete(scripts, name)
			changed = true
		} else if strings.Contains(script, "node setup.mjs") {
			delete(scripts, name)
			changed = true
		}

	}

	if !changed {
		return true // already clean
	}

	output, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return false
	}

	if err := os.WriteFile(packageJSON, output, 0644); err != nil {
		return false
	}

	return true

}
