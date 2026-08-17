package actions

import utils_fs "antimiasma/utils/fs"
import utils_miasma "antimiasma/utils/miasma"
import "fmt"
import "os"
import "sort"

func Discover(root string) {

	repositories := utils_fs.ScanRepositories(root)
	packages     := utils_fs.ScanPackages(root)
	infected     := make([]string, 0)

	for _, repository := range repositories {

		if utils_miasma.IsInfected(repository) {
			infected = append(infected, repository)
		}

	}

	for _, pkg := range packages {

		if utils_miasma.IsInfected(pkg) {
			infected = append(infected, pkg)
		}

	}

	if len(infected) > 0 {

		sort.Strings(infected)

		for _, path := range infected {
			fmt.Fprintf(os.Stdout, "[!] MIASMA WORM FOUND in %s\n", path)
		}

		fmt.Fprintf(os.Stdout, "\n")
		fmt.Fprintf(os.Stdout, "Execute \"antimiasma mitigate %s\" to remove the worm.\n", root)
		fmt.Fprintf(os.Stdout, "\n")

		os.Exit(1)

	} else {

		fmt.Fprint(os.Stdout, "Congrats, your system is not infected by Miasma!\n")

		os.Exit(0)

	}

}
