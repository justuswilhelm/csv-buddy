package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/justuswilhelm/csv-buddy/lib"
	"os"
)

// Config stores flags
type Config struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
}

var (
	// Config store the application configuration
	config = Config{}

	// Parser parses args into application configuration
	parser = flags.NewParser(&config, flags.Default)
)

func init() {
	var summaryCommand lib.SummaryCommand
	parser.AddCommand(
		"summary",
		"Print Summary",
		"The summary command prints column based summaries for a given file.",
		&summaryCommand)
}

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
