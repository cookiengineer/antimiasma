package miasma

import "os"
import "path/filepath"

func HasImplant(repo string) bool {

	implantPath := filepath.Join(repo, ".github", "setup.js")

	info, err := os.Stat(implantPath)
	if err != nil {
		return false
	}

	return !info.IsDir()

}
