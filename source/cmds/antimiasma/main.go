package main

import "antimiasma/actions"
import utils_cli "antimiasma/utils/cli"
import "os"
import "path/filepath"
import "strings"

func main() {

	command := ""
	root := ""

	if len(os.Args) == 3 {

		tmp1 := strings.ToLower(strings.TrimSpace(os.Args[1]))

		if tmp1 == "discover" {
			command = "discover"
		} else if tmp1 == "mitigate" {
			command = "mitigate"
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
		} else if tmp1 == "mitigate" {
			command = "mitigate"
		}

		cwd, err := os.Getwd()

		if err == nil {

			volume := filepath.VolumeName(cwd)

			if volume == "" {
				// Unix
				root = string(filepath.Separator)
			} else {
				// Windows
				root = volume + string(filepath.Separator)
			}

		}

	}

	if command != "" && root != "" {

		if command == "discover" {
			actions.Discover(root)
		} else if command == "mitigate" {
			actions.Mitigate(root)
		}

	} else {

		utils_cli.ShowUsage()
		os.Exit(1)

	}

}
