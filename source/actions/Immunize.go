package actions

import utils_antimiasma "antimiasma/utils/antimiasma"
import utils_fs "antimiasma/utils/fs"
import utils_miasma "antimiasma/utils/miasma"
import "fmt"
import "os"
import "path/filepath"
import "sort"

func Immunize(root string, self string) {

	repositories := utils_fs.ScanRepositories(root)
	packages     := utils_fs.ScanPackages(root)
	infected     := make([]string, 0)

	for _, repository := range repositories {

		if utils_miasma.IsInfected(repository) {

			if repository != self && utils_antimiasma.IsVaccinated(repository) == false {
				infected = append(infected, repository)
			}

		}

	}

	for _, pkg := range packages {

		if utils_miasma.IsInfected(pkg) {

			if pkg != self && utils_antimiasma.IsVaccinated(pkg) == false {
				infected = append(infected, pkg)
			}

		}

	}

	if len(infected) > 0 {

		sort.Strings(infected)

		messages     := make([]string, 0)
		vaccinated   := make([]string, 0)
		unvaccinated := make([]string, 0)

		for _, repository := range infected {

			is_vaccinated := utils_antimiasma.Vaccinate(repository)

			if is_vaccinated == true {
				vaccinated = append(vaccinated, repository)
				messages = append(messages, fmt.Sprintf("- Miasma worm vaccinated in %s\n", repository))
			} else {
				unvaccinated = append(unvaccinated, repository)
				messages = append(messages, fmt.Sprintf("- [!] MIASMA WORM REMAINED in %s\n", repository))
			}

		}

		report_path := filepath.Join(self, "README-IMPORTANT-ANTIMIASMA.md")
		report_bytes := utils_antimiasma.CreateReport(report_path, messages)

		err := os.WriteFile(report_path, report_bytes, 0666)

		if err == nil {

			fmt.Fprintf(os.Stdout, "\n")
			fmt.Fprintf(os.Stdout, "Miasma Vaccination Report has been written to %s for the developer to discover!\n", report_path)
			fmt.Fprintf(os.Stdout, "\n")

		}

		if len(unvaccinated) > 0 {

			fmt.Fprintf(os.Stdout, "\n")
			fmt.Fprintf(os.Stdout, "!! Some Miasma worms remained and couldn't be removed!\n")
			fmt.Fprintf(os.Stdout, "!! Immediately remove the worm implants manually!\n")
			fmt.Fprintf(os.Stdout, "!! Don't open ANY AI assisted IDE to do that!\n")
			fmt.Fprintf(os.Stdout, "\n")

			CheckSelfDestroy(self, len(vaccinated), len(unvaccinated))

			os.Exit(1)

		} else {

			fmt.Fprintf(os.Stdout, "\n")
			fmt.Fprintf(os.Stdout, "All discovered Miasma worms have been vaccinated.\n")
			fmt.Fprintf(os.Stdout, "Have a great day!\n")
			fmt.Fprintf(os.Stdout, "\n")

			CheckSelfDestroy(self, len(vaccinated), len(unvaccinated))

			os.Exit(0)

		}

	} else {

		fmt.Fprint(os.Stdout, "Congrats, your system is not infected by Miasma!\n")

		CheckSelfDestroy(self, 0, 0)

		os.Exit(0)

	}

}
