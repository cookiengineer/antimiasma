package miasma

import utils_json "antimiasma/utils/json"
import "os"
import "path/filepath"

func FixComposer(repo string) bool {

	composer_path := filepath.Join(repo, "composer.json")

	data, err := os.ReadFile(composer_path)
	if err != nil {
		return true // nothing to fix
	}

	var pkg map[string]any
	if err := utils_json.Unmarshal(data, &pkg); err != nil {
		return false
	}

	scripts, ok := pkg["scripts"].(map[string]any)
	if !ok {
		return true // nothing to fix
	}

	cleaned, changed := utils_json.FilterVal(scripts, []string{
		"node .github/setup.js",
		"node setup.mjs",
	})

	if !changed {
		return true // already clean
	}

	pkg["scripts"] = cleaned

	output, err := utils_json.Marshal(pkg)
	if err != nil {
		return false
	}

	if err := os.WriteFile(composer_path, output, 0644); err != nil {
		return false
	}

	return true
}
