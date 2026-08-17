package miasma

import "os"
import "path/filepath"
import "strings"

func IsCursorAffected(repo string) bool {

	rules_path := filepath.Join(repo, ".cursor", "rules", "setup.mdc")

	data, err := os.ReadFile(rules_path)
	if err != nil {
		return false
	}

	return strings.Contains(string(data), "node .github/setup.js")

}
