package main

import "antimiasma/actions"
import utils_cli "antimiasma/utils/cli"
import utils_fs "antimiasma/utils/fs"
import "os"
import "path/filepath"
import "strings"

func main() {

	command := ""
	root := ""
	self := ""

	if len(os.Args) == 3 {

		tmp1 := strings.ToLower(strings.TrimSpace(os.Args[1]))

		if tmp1 == "discover" {
			command = "discover"
		} else if tmp1 == "mitigate" {
			command = "mitigate"
		} else if tmp1 == "vaccinate" {
			command = "vaccinate"
		}

		tmp2 := strings.TrimSpace(os.Args[2])
		tmp3, err1 := filepath.Abs(tmp2)

		if err1 == nil {

			stat, err2 := os.Stat(tmp3)

			if err2 == nil && stat.IsDir() {
				root = tmp3
			}

		}

	} else if len(os.Args) == 2 {

		tmp1 := strings.ToLower(strings.TrimSpace(os.Args[1]))

		if tmp1 == "discover" {
			command = "discover"
		} else if tmp1 == "immunize" {
			command = "immunize"
		} else if tmp1 == "mitigate" {
			command = "mitigate"
		} else if tmp1 == "vaccinate" {
			command = "vaccinate"
		}

		cwd, err := os.Getwd()

		if err == nil {

			volume := filepath.VolumeName(cwd)

			if volume == "" {
				// Unix
				root = string(filepath.Separator)

				repo, err2 := utils_fs.ParentRepository(self)

				if err2 == nil {
					self = repo
				}

			} else {
				// Windows
				root = volume + string(filepath.Separator)

				repo, err2 := utils_fs.ParentRepository(self)

				if err2 == nil {
					self = repo
				}

			}

		}

	}

	if command != "" && root != "" {

		if command == "discover" {
			actions.Discover(root)
		} else if command == "immunize" {
			actions.Immunize(root, self)
		} else if command == "mitigate" {
			actions.Mitigate(root)
		} else if command == "vaccinate" {
			actions.Vaccinate(root)
		}

	} else {

		utils_cli.ShowVaccineUsage()
		os.Exit(1)

	}

}
