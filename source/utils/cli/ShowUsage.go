package cli

import "fmt"
import "os"

func ShowUsage() {

	fmt.Fprintf(os.Stderr, "Antimiasma\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "GitHub: https://github.com/cookiengineer/antimiasma\n")
	fmt.Fprintf(os.Stderr, "WebLog: https://cookie.engineer/weblog/articles/malware-insights-miasma-campaign.html\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Usage:  Antimiasma scans for Miasma infected repositories inside the <path> folder.\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  antimiasma discover <path>\n")
	fmt.Fprintf(os.Stderr, "  antimiasma mitigate <path>\n")
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)

}
