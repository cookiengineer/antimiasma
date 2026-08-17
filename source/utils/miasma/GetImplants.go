package miasma

import "os"
import "path/filepath"

func GetImplants(repo string) []string {

	files := []string{
		".github/setup.js",
		"_index.js",
		"src/hooks/deps",
		".claude/math_init.js",
		".claude/setup.mjs",
		".gemini/math_init.js",
		".gemini/setup.mjs",
		".vscode/math_init.js",
		".vscode/setup.mjs",
		"setup.mjs",
		"Math_Symbol.js",
	}

	result := make([]string, 0)

	for _, file := range files {

		stat, err := os.Stat(filepath.Join(repo, file))

		if err == nil && !stat.IsDir() {
			result = append(result, file)
		}

	}

	return result

}
