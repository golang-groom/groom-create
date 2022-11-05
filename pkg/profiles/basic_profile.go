package profile

import (
	_ "embed"
	"path/filepath"

	"github.com/pspiagicw/groom-create/pkg/log"
)

// Template for the Makefile
//
//go:embed data/default/makefile
var makeFileTemplateStringDefault string

// Template for the main file
//
//go:embed data/default/mainfile
var mainFileTemplateStringDefault string

// Instance of a `ProfileCreator`.
// Creates a Basic Profile
type BasicProfileCreator struct {
	Template *ProjectTemplate
	Log      log.Logger
}

func (d *BasicProfileCreator) GetProjectName() string {
	return d.Template.ProjName
}

func (d *BasicProfileCreator) GetProjectURL() string {
	return d.Template.ProjURL
}

// Steps for creating a basic project
func (d *BasicProfileCreator) CreateProfile() {

	d.createMakeFile()
	d.createMainFile()
}

func (d *BasicProfileCreator) createMainFile() {
	contents := d.getMainFileContents()

	mainFilePath := filepath.Join(d.GetProjectName(), "main.go")

	writeFile(d.Log, mainFilePath, contents)

}

func (d *BasicProfileCreator) createMakeFile() {
	contents := d.getMakefileContents()

	makeFilePath := filepath.Join(d.GetProjectName(), "Makefile")

	writeFile(d.Log, makeFilePath, contents)

}

func (d *BasicProfileCreator) getMakefileContents() string {

	return populateTemplate(d.Log, makeFileTemplateStringDefault, d.Template)
}

func (d *BasicProfileCreator) getMainFileContents() string {
	return mainFileTemplateStringDefault
}
