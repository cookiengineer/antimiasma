package miasma

import utils_json "antimiasma/utils/json"
import "os"
import "path/filepath"

func IsNodeAffected(repo string) bool {

	packageJSON := filepath.Join(repo, "package.json")

	data, err := os.ReadFile(packageJSON)
	if err != nil {
		return false
	}

	var pkg map[string]any
	if err := utils_json.Unmarshal(data, &pkg); err != nil {
		return false
	}

	scripts, ok := pkg["scripts"].(map[string]any)
	if !ok {
		return false
	}

	return utils_json.ContainsVal(scripts, []string{
		"node .github/setup.js",
		"src/hooks/deps",
		"node setup.mjs",
	})

}
