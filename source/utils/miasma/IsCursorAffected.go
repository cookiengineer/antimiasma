package miasma

import "antimiasma/utils/log"
import "os"
import "path/filepath"
import "strings"

func IsCursorAffected(repo string) bool {

	rulesPath := filepath.Join(repo, ".cursor", "rules", "setup.mdc")

	log.Printf("  checking Cursor: %s\n", rulesPath)

	data, err := os.ReadFile(rulesPath)
	if err != nil {
		return false
	}

	return strings.Contains(string(data), "node .github/setup.js")

}
