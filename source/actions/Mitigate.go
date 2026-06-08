package actions

import utils_fs "antimiasma/utils/fs"
import utils_miasma "antimiasma/utils/miasma"
import "fmt"
import "os"
import "sort"

func Mitigate(root string) {

	repositories := utils_fs.ScanRepositories(root)
	infected := make([]string, 0)

	for _, repository := range repositories {

		if utils_miasma.IsInfected(repository) {
			infected = append(infected, repository)
		}

	}

	if len(infected) > 0 {

		sort.Strings(infected)

		unfixed := make([]string, 0)

		for _, repository := range infected {

			fixed := utils_miasma.Fix(repository)

			if fixed == true {
				fmt.Fprintf(os.Stdout, "    Miasma worm removed from %s\n", repository)
			} else {
				unfixed = append(unfixed, repository)
				fmt.Fprintf(os.Stdout, "[!] MIASMA WORM REMAINED in %s\n", repository)
			}

		}

		if len(unfixed) > 0 {

			fmt.Fprintf(os.Stdout, "\n")
			fmt.Fprintf(os.Stdout, "!! Some Miasma worms remained and couldn't be removed!\n")
			fmt.Fprintf(os.Stdout, "!! Immediately remove the worm implants manually!\n")
			fmt.Fprintf(os.Stdout, "!! Don't open ANY AI assisted IDE to do that!\n")
			fmt.Fprintf(os.Stdout, "\n")

			os.Exit(1)

		} else {

			fmt.Fprintf(os.Stdout, "\n")
			fmt.Fprintf(os.Stdout, "All discovered Miasma worms have been removed.\n")
			fmt.Fprintf(os.Stdout, "Have a great day!\n")
			fmt.Fprintf(os.Stdout, "\n")

			os.Exit(0)

		}

	} else {

		fmt.Fprint(os.Stdout, "Congrats, your system is not infected by Miasma!\n")

		os.Exit(0)

	}

}
