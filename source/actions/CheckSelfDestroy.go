package actions

import "os"
import "path/filepath"
import "time"

func CheckSelfDestroy(repo string, vaccinated int, unvaccinated int) {

	if vaccinated == 0 && unvaccinated == 0 {

		exe, err0 := os.Executable()

		if err0 == nil {

			real_path, err1 := filepath.EvalSymlinks(exe)

			if err1 == nil {

				stat, err2 := os.Stat(real_path)

				if err2 == nil {

					if time.Since(stat.ModTime()) > 7 * 24 * time.Hour {

						Mitigate(repo)

						os.Remove(real_path)

					}

				}

			}

		}

	}

}
