package profile

import (
	_ "embed"
	"path/filepath"

	"github.com/pspiagicw/colorlog"
)

//go:embed data/lib/libfile
var libFileTemplateLIB string

type LIBProfileCreator struct {
    Template *ProjectTemplate
    Log colorlog.ColorLogger
}

func (l *LIBProfileCreator) GetProjectName() string {
    return l.Template.ProjName
}

func (l *LIBProfileCreator) GetProjectURL() string {
    return l.Template.ProjURL
}

func (l *LIBProfileCreator) CreateProfile() {

    l.createLibFile()
}

func (l *LIBProfileCreator) createLibFile() {
    contents := l.getLibFileContents()

    libFilePath := filepath.Join(l.GetProjectName(), "lib.go")

    writeFile(l.Log, libFilePath, contents)
}

func (l *LIBProfileCreator) getLibFileContents() string {
    return populateTemplate(l.Log, libFileTemplateLIB, l.Template)
}
