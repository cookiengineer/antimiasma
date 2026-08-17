package cli

import "fmt"
import "os"

func ShowVaccineUsage() {

	fmt.Fprintf(os.Stderr, "Antimiasma\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "GitHub: https://github.com/cookiengineer/antimiasma\n")
	fmt.Fprintf(os.Stderr, "WebLog: https://cookie.engineer/weblog/articles/malware-insights-miasma-campaign.html\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Usage:  Antimiasma scans for Miasma infected repositories inside the <path> folder.\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  antimiasma discover <path>\n")
	fmt.Fprintf(os.Stderr, "  antimiasma mitigate <path>\n")
	fmt.Fprintf(os.Stderr, "  antimiasma vaccinate <path>\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "IMPORTANT:\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  The vaccinate action will implant this program and spread it as an anti-worm\n")
	fmt.Fprintf(os.Stderr, "  vaccine using the same spreading technique to fix infected machines.\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  It will delete itself after a timeout of 7 days, when no further infected\n")
	fmt.Fprintf(os.Stderr, "  repositories were discovered on that machine.\n")
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)

}

