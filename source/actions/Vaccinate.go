package actions

import utils_antimiasma "antimiasma/utils/antimiasma"
import utils_fs "antimiasma/utils/fs"
import utils_miasma "antimiasma/utils/miasma"
import "fmt"
import "os"
import "sort"

func Vaccinate(root string) {

	repositories := utils_fs.ScanRepositories(root)
	packages     := utils_fs.ScanPackages(root)
	infected     := make([]string, 0)

	for _, repository := range repositories {

		if utils_miasma.IsInfected(repository) {

			if utils_antimiasma.IsVaccinated(repository) == false {
				infected = append(infected, repository)
			}

		}

	}

	for _, pkg := range packages {

		if utils_miasma.IsInfected(pkg) {

			if utils_antimiasma.IsVaccinated(pkg) == false {
				infected = append(infected, pkg)
			}

		}

	}

	if len(infected) > 0 {

		sort.Strings(infected)

		unvaccinated := make([]string, 0)

		for _, path := range infected {

			vaccinated := utils_antimiasma.Vaccinate(path)

			if vaccinated == true {
				fmt.Fprintf(os.Stdout, "    Miasma worm vaccinated in %s\n", path)
			} else {
				unvaccinated = append(unvaccinated, path)
				fmt.Fprintf(os.Stdout, "[!] MIASMA WORM REMAINED in %s\n", path)
			}

		}

		if len(unvaccinated) > 0 {

			fmt.Fprintf(os.Stdout, "\n")
			fmt.Fprintf(os.Stdout, "!! Some Miasma worms remained and couldn't be removed!\n")
			fmt.Fprintf(os.Stdout, "!! Immediately remove the worm implants manually!\n")
			fmt.Fprintf(os.Stdout, "!! Don't open ANY AI assisted IDE to do that!\n")
			fmt.Fprintf(os.Stdout, "\n")

			os.Exit(1)

		} else {

			fmt.Fprintf(os.Stdout, "\n")
			fmt.Fprintf(os.Stdout, "All discovered Miasma worms have been vaccinated.\n")
			fmt.Fprintf(os.Stdout, "Have a great day!\n")
			fmt.Fprintf(os.Stdout, "\n")

			os.Exit(0)

		}

	} else {

		fmt.Fprint(os.Stdout, "Congrats, your system is not infected by Miasma!\n")

		os.Exit(0)

	}

}

