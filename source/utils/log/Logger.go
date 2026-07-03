package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

var writer io.Writer = os.Stderr

func Printf(format string, args ...any) {
	timestamp := time.Now().Format("15:04:05.000")
	fmt.Fprintf(writer, "[%s] %s", timestamp, fmt.Sprintf(format, args...))
}

func Println(args ...any) {
	timestamp := time.Now().Format("15:04:05.000")
	fmt.Fprintf(writer, "[%s] ", timestamp)
	fmt.Fprintln(writer, args...)
}

func Fprintf(w io.Writer, format string, args ...any) {
	timestamp := time.Now().Format("15:04:05.000")
	fmt.Fprintf(w, "[%s] %s", timestamp, fmt.Sprintf(format, args...))
}
