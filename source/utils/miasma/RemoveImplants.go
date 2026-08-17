package miasma

import utils_elf "antimiasma/utils/elf"
import utils_js "antimiasma/utils/js"
import "os"
import "path/filepath"

func RemoveImplants(repo string) bool {

	errors := int(0)

	binaries := []string{
		"src/hooks/deps",
	}

	for _, binary := range binaries {

		data, err1 := os.ReadFile(filepath.Join(repo, binary))

		if err1 == nil && utils_elf.IsSuspicious(data) {

			err2 := os.Remove(filepath.Join(repo, binary))

			if err2 != nil {
				errors++
			}

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

		data, err1 := os.ReadFile(filepath.Join(repo, script))

		if err1 == nil && utils_js.IsSuspicious(data) {

			err2 := os.Remove(filepath.Join(repo, script))

			if err2 != nil {
				errors++
			}

		}

	}

	if errors > 0 {
		return false
	} else {
		return true
	}

}

