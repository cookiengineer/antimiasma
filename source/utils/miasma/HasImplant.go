package miasma

import utils_elf "antimiasma/utils/elf"
import utils_js "antimiasma/utils/js"
import "os"
import "path/filepath"

func HasImplant(repo string) bool {

	result := false

	binaries := []string{
		"src/hooks/deps",
	}

	for _, binary := range binaries {

		data, err := os.ReadFile(filepath.Join(repo, binary))

		if err == nil && utils_elf.IsSuspicious(data) {
			result = true
			break
		}

	}

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
			result = true
			break
		}

	}

	return result

}
