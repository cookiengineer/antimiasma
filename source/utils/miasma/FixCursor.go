package miasma

import "os"
import "path/filepath"

func FixCursor(repo string) bool {

	rulesPath := filepath.Join(repo, ".cursor", "rules", "setup.mdc")

	err := os.Remove(rulesPath)
	if err == nil {
		return true
	}

	// If it's already gone, consider it fixed
	if os.IsNotExist(err) {
		return true
	}

	return false

}
