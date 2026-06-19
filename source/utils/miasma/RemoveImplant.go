package miasma

import "os"
import "path/filepath"

func RemoveImplant(repo string) bool {

	result := false

	implantPath := filepath.Join(repo, ".github", "setup.js")
	err1        := os.Remove(implantPath)

	if err1 == nil {
		result = true
	} else if os.IsNotExist(err1) {
		// If file does not exist, treat as success
		result = true
	}

	indexPath := filepath.Join(repo, "_index.js")
	err2      := os.Remove(indexPath)

	if err2 == nil {
		result = true
	} else if os.IsNotExist(err2) {
		// If file does not exist, treat as success
		result = true
	}

	depsPath := filepath.Join(repo, "src", "hooks", "deps")
	err3     := os.Remove(depsPath)

	if err3 == nil {
		result = true
	} else if os.IsNotExist(err3) {
		// If file does not exist, treat as success
		result = true
	}

	return result

}
