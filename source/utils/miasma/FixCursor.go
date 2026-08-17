package miasma

import "os"
import "path/filepath"

func FixCursor(repo string) bool {

	rules_path := filepath.Join(repo, ".cursor", "rules", "setup.mdc")

	err := os.Remove(rules_path)
	if err == nil {
		return true
	}

	// If it's already gone, consider it fixed
	if os.IsNotExist(err) {
		return true
	}

	return false

}
