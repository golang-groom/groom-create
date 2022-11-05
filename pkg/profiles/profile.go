// Package to contain all profiles.
package profile

import (
	"bytes"
	"github.com/pspiagicw/groom-create/pkg/log"
	"os"
	"text/template"
)

/*
Profile Creator is an inteface which defines all type of profiles.

Dependencies:
- Project Name
- Project URL

Must implement:
  - CreateProfile(): This is a stepwise creation of all the files and folders.
    Every given profile must statisfy this function..
*/
type ProfileCreator interface {
	CreateProfile()
	GetProjectName() string
	GetProjectURL() string
}

// Template for Project Info
type ProjectTemplate struct {
	ProjName string
	ProjURL  string
}

// Helper Functions to write a file.
// Automatically logs result.
func writeFile(log log.Logger, filename string, contents string) {
	err := os.WriteFile(filename, []byte(contents), 0644)

	if err != nil {
		log.LogFatalf("Cannot create '%s' , %q", filename, err)
	} else {
		log.LogSuccess("Successfully written '%s'", filename)
	}

}

// Helper function to populate Makefiles.
// Takes MakefileTemplate and Template String.
// Logs the result
func populateTemplate(log log.Logger, templateString string, generator *ProjectTemplate) string {
	makefileTemplate, err := template.New("Makefile").Parse(templateString)
	if err != nil {
		log.LogFatalf("Failed to parse template %q", err)
	}

	var output bytes.Buffer

	err = makefileTemplate.Execute(&output, generator)

	if err != nil {
		log.LogFatalf("Unable to execute template %q", err)
	}

	return output.String()
}
