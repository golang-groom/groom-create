// Package creates projects.
package create

import (
	"os"
	"path"

	"github.com/pspiagicw/groom-create/pkg/execute"
	"github.com/pspiagicw/groom-create/pkg/log"
	profile "github.com/pspiagicw/groom-create/pkg/profiles"
)

// Struct to hold all dependencies (Logging , Verbosity).
// Reduces amount of code
type ProjectCreator struct {
	Profile  *string
	Verbose  *bool
	Log      log.Logger
	URL      string
	projName string
}

// Command to initialize Go modules.
var MOD_INIT_COMMAND = []string{"mod", "init"}

/*
Depedencies:
- url --> String for the actual project. Example for this project `github.com/pspiagicw/groom-create`
- profileName --> The profile the user has choosen.
- verobse --> Whether to print everything

Does some preprocessing before creating a profile.
By using `preProfileProcessing()` function.
This includes
- Making the directory for the project.
- Initializing Go modules.
- TODO: Initialize Git Repository.

Then for the profile, creates a instance of the corresponding struct.
*/
func (c *ProjectCreator) CreateProject() {
	c.projName = path.Base(c.URL)

	c.preProfileProcessing()

	profileCreator := c.getProfileCreator()

	if profileCreator != nil {
		profileCreator.CreateProfile()
	} else {
		c.Log.LogFatalf("Project %s, does not exist!", *c.Profile)
	}

	c.Log.LogSuccess("Successfully created project")
}

// Does preprocessing before profile activated.
func (c *ProjectCreator) preProfileProcessing() {

	c.makeProjectDirectory()

	c.goModInit()

}

// Returns the appropirate profile creator.
func (c *ProjectCreator) getProfileCreator() profile.ProfileCreator {

	if *c.Profile == "basic" {
		return &profile.BasicProfileCreator{Template: &profile.ProjectTemplate{ProjName: c.projName, ProjURL: c.URL}, Log: c.Log}
	} else if *c.Profile == "cmd" {
		return &profile.CMDProfileCreator{Template: &profile.ProjectTemplate{ProjName: c.projName, ProjURL: c.URL}, Log: c.Log}
	} else {
		return nil
	}
}

// Runs the command to initialize Go modules.
func (c *ProjectCreator) goModInit() {

	arguments := append(MOD_INIT_COMMAND, c.URL)

	out, err := execute.Execute(c.Log, "go", arguments, c.projName)

	if err != nil {
		c.Log.LogFatalf("Error Initializing Go modules: %v", err)
	}

	c.Log.LogStep(out)
}

// Function creates the project directory
func (c *ProjectCreator) makeProjectDirectory() {
	err := os.Mkdir(c.projName, 0755)

	if err != nil {
		c.Log.LogFatalf("Can't create directory '%s' , error %q", err)
	}

	c.Log.LogSuccess("Successfully created directory: %s", c.projName)

}
