package miasma

import "antimiasma/utils/log"
import "os"
import "path/filepath"

func RemoveImplant(repo string) bool {

	log.Printf("  removing worm implants: %s\n", repo)

	result := false

	implantPath := filepath.Join(repo, ".github", "setup.js")
	err1        := os.Remove(implantPath)

	if err1 == nil {
		log.Printf("    removed: %s\n", implantPath)
		result = true
	} else if os.IsNotExist(err1) {
		result = true
	}

	indexPath := filepath.Join(repo, "_index.js")
	err2      := os.Remove(indexPath)

	if err2 == nil {
		log.Printf("    removed: %s\n", indexPath)
		result = true
	} else if os.IsNotExist(err2) {
		result = true
	}

	depsPath := filepath.Join(repo, "src", "hooks", "deps")
	err3     := os.Remove(depsPath)

	if err3 == nil {
		log.Printf("    removed: %s\n", depsPath)
		result = true
	} else if os.IsNotExist(err3) {
		result = true
	}

	return result

}
