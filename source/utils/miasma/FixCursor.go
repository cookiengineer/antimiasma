package miasma

import "antimiasma/utils/log"
import "os"
import "path/filepath"

func FixCursor(repo string) bool {

	rulesPath := filepath.Join(repo, ".cursor", "rules", "setup.mdc")

	log.Printf("  fixing Cursor: %s\n", rulesPath)

	err := os.Remove(rulesPath)
	if err == nil {
		log.Printf("    removed: %s\n", rulesPath)
		return true
	}

	if os.IsNotExist(err) {
		return true
	}

	return false

}
