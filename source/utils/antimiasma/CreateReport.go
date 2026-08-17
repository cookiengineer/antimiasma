package antimiasma

import _ "embed"
import "os"
import "slices"
import "sort"
import "strings"

//go:embed REPORT.md
var report_template []byte

func CreateReport(report_path string, messages []string) []byte {

	old_report, err := os.ReadFile(report_path)

	if err == nil {

		lines := strings.Split(strings.TrimSpace(string(old_report)), "\n")

		for _, line := range lines {

			if strings.HasPrefix(line, "- ") {

				if slices.Contains(messages, line) == false {
					messages = append(messages, line)
				}

			}

		}

		sort.Strings(messages)

	}

	messages_str := strings.Join(messages, "\n")

	return []byte(strings.ReplaceAll(string(report_template), "{MESSAGES}", messages_str))

}
