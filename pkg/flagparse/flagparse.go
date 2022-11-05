// Package to parse flags
package flagparse

import (
	"flag"
	"fmt"
	"os"

	"github.com/pspiagicw/groom-create/pkg/create"
	"github.com/pspiagicw/groom-create/pkg/log"
)

// Function to define flags and parse them
func ParseAndRun() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		fmt.Fprintf(os.Stderr, "groom [OPTIONS] [NAME] \n")

		flag.PrintDefaults()

	}

	verbose := flag.Bool("verbose", false, "Whether to log verbosely")
	profile := flag.String("profile", "basic", "Profile to use while creating a project")

	flag.Parse()

	// Default Logger
	logger := &log.DefaultLogger{}

	url := flag.Arg(0)
	if url != "" {
		creator := &create.ProjectCreator{
			Profile: profile,
			Verbose: verbose,
			Log:     logger,
			URL:     url,
		}
		creator.CreateProject()
	} else {
		flag.Usage()
		logger.LogFatalf("Empty project name provided")
	}
}
