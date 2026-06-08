package miasma

import "os"
import "path/filepath"

func RemoveImplant(repo string) bool {

	implantPath := filepath.Join(repo, ".github", "setup.js")

	err := os.Remove(implantPath)
	if err == nil {
		return true
	}

	// If file does not exist, treat as success
	if os.IsNotExist(err) {
		return true
	}

	return false

}
