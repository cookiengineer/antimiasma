package miasma

import utils_js "antimiasma/utils/js"
import "os"
import "path/filepath"

func GetImplants(repo string) []string {

	result := make([]string, 0)

	scripts := []string{
		".github/setup.js",
		"_index.js",
		".claude/math_init.js",
		".claude/setup.mjs",
		".gemini/math_init.js",
		".gemini/setup.mjs",
		".vscode/math_init.js",
		".vscode/setup.mjs",
		"setup.mjs",
		"Math_Symbol.js",
	}

	for _, script := range scripts {

		data, err := os.ReadFile(filepath.Join(repo, script))

		if err == nil && utils_js.IsSuspicious(data) {
			result = append(result, script)
		}

	}

	return result

}
