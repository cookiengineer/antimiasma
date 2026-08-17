package miasma

import utils_json "antimiasma/utils/json"
import "os"
import "path/filepath"

func IsComposerAffected(repo string) bool {

	composer_path := filepath.Join(repo, "composer.json")

	data, err := os.ReadFile(composer_path)
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
		"node setup.mjs",
	})

}
