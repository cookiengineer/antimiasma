package miasma

import "os"
import "path/filepath"
import "strings"

func IsCursorAffected(repo string) bool {

	rulesPath := filepath.Join(repo, ".cursor", "rules", "setup.mdc")

	data, err := os.ReadFile(rulesPath)
	if err != nil {
		return false
	}

	return strings.Contains(string(data), "node .github/setup.js")

}
