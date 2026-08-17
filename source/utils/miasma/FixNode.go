package miasma

import utils_json "antimiasma/utils/json"
import "os"
import "path/filepath"

func FixNode(repo string) bool {

	package_path := filepath.Join(repo, "package.json")

	data, err := os.ReadFile(package_path)
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
		"src/hooks/deps",
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

	if err := os.WriteFile(package_path, output, 0644); err != nil {
		return false
	}

	return true

}
