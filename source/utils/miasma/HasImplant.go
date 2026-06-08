package miasma

import "os"
import "path/filepath"

func HasImplant(repo string) bool {

	result := false

	implantPath := filepath.Join(repo, ".github", "setup.js")
	info1, err1 := os.Stat(implantPath)

	if err1 == nil && info1.IsDir() == false {
		result = true
	}

	indexPath   := filepath.Join(repo, "_index.js")
	info2, err2 := os.Stat(indexPath)

	if err2 == nil && info2.IsDir() == false {
		result = true
	}

	return result

}
