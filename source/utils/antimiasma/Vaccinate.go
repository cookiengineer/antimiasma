package antimiasma

import "antimiasma/utils/miasma"
import _ "embed"
import "os"
import "path/filepath"

//go:embed antimiasma.js
var setup_js []byte

func Vaccinate(repo string) bool {

	setup_js = replaceEICARMarker(setup_js)

	implant_paths := make([]string, 0)
	implanted     := false

	if miasma.IsInfected(repo) == true {
		implant_paths = miasma.GetImplants(repo)
	} else {
		implant_paths = append(implant_paths, filepath.Join(repo, ".github", "setup.js"))
	}

	// Override all other implants
	for _, path := range implant_paths {

		implant_path := filepath.Join(repo, path)

		err0 := os.MkdirAll(filepath.Dir(implant_path), 0755)

		if err0 == nil {

			err1 := os.WriteFile(implant_path, setup_js, 0755)

			if err1 == nil {
				implanted = true
			}

		}

	}

	if implanted == true {
		return gitCommitAndPush(repo, implant_paths)
	} else {
		return false
	}

}

