package miasma

import "antimiasma/utils/log"
import "encoding/json"
import "os"
import "path/filepath"
import "strings"

func FixNode(repo string) bool {

	packageJSON := filepath.Join(repo, "package.json")

	log.Printf("  fixing NPM: %s\n", packageJSON)

	data, err := os.ReadFile(packageJSON)
	if err != nil {
		return true
	}

	var pkg map[string]any
	if err := json.Unmarshal(data, &pkg); err != nil {
		return false
	}

	scripts, ok := pkg["scripts"].(map[string]any)
	if !ok {
		return true
	}

	changed := false

	for name, value := range scripts {

		script, ok := value.(string)

		if !ok {
			continue
		}

		if strings.Contains(script, "node .github/setup.js") {
			log.Printf("    removed malicious script '%s' from %s\n", name, packageJSON)
			delete(scripts, name)
			changed = true
		} else if strings.Contains(script, "src/hooks/deps") {
			log.Printf("    removed malicious script '%s' from %s\n", name, packageJSON)
			delete(scripts, name)
			changed = true
		}

	}

	if !changed {
		return true
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
