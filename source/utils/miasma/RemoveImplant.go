package miasma

import "os"
import "path/filepath"

func removeIfExists(path string) bool {

	err := os.Remove(path)

	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		// If file does not exist, treat as success
		return true
	}

	return false

}

func RemoveImplant(repo string) bool {

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

	result := false

	for _, file := range files {

		if removeIfExists(filepath.Join(repo, file)) {
			result = true
		}

	}

	return result

}

