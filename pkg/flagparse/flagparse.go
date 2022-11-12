// Package to parse flags
package flagparse

import (
	"flag"
	"fmt"
	"os"

	"github.com/pspiagicw/colorlog"
	"github.com/pspiagicw/groom-create/pkg/create"
)

// Function to define flags and parse them
func ParseAndRun() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		fmt.Fprintf(os.Stderr, "groom create [OPTIONS] [NAME] \n")

		flag.PrintDefaults()

	}

	profile := flag.String("profile", "basic", "Profile to use while creating a project")

	flag.Parse()

	// Default Logger
	logger := &colorlog.DefaultColorLogger{}

	url := flag.Arg(0)
	if url != "" {
		creator := &create.ProjectCreator{
			Profile: profile,
			Verbose: false,
			Log:     logger,
			URL:     url,
		}
		creator.CreateProject()
	} else {
		flag.Usage()
		logger.LogFatal("Empty project name provided")
	}
}
