package antimiasma

import "antimiasma/utils/miasma"
import "os"
import "path/filepath"

func IsVaccinated(repo string) bool {

	implant_paths := miasma.GetImplants(repo)

	if len(implant_paths) > 0 {

		result := true

		for _, implant_path := range implant_paths {

			data, err := os.ReadFile(filepath.Join(repo, implant_path))

			if err == nil {

				if containsEICARMarker(string(data)) {
					// Ignore
				} else {
					result = false
				}

			} else {
				result = false
			}

		}

		return result

	} else {
		// No implants found
		return true
	}

}
