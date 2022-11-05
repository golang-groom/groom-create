package profile

import (
	_ "embed"
	"github.com/pspiagicw/groom-create/pkg/log"
	"os"
	"path/filepath"
)

//go:embed data/cmd/pkgfile
var pkgFileTemplateCMD string

//go:embed data/cmd/mainfile
var mainFileTemplateStringCMD string

//go:embed data/cmd/makefile
var makeFileTemplateStringCMD string

type CMDProfileCreator struct {
	Template *ProjectTemplate
	Log      log.Logger
}

func (c *CMDProfileCreator) GetProjectName() string {
	return c.Template.ProjName
}

func (c *CMDProfileCreator) GetProjectURL() string {
	return c.Template.ProjURL
}

func (c *CMDProfileCreator) CreateProfile() {

	c.createMakeFile()
	c.createMainFile()
	c.createPkgFile()
}

func (c *CMDProfileCreator) createPkgFile() {
	contents := c.getPkgFileContents()

	pkgFilePath := filepath.Join(c.GetProjectName(), "pkg")
	pkgFilePath = filepath.Join(pkgFilePath, "simple")

	err := os.MkdirAll(pkgFilePath, 0755)
	if err != nil {
		c.Log.LogFatalf("Error creating directory '%s' and it's subdirectories", pkgFilePath)
	}

	pkgFilePath = filepath.Join(pkgFilePath, "simple.go")
	writeFile(c.Log, pkgFilePath, contents)
}

func (c *CMDProfileCreator) getPkgFileContents() string {

	return populateTemplate(c.Log, pkgFileTemplateCMD, c.Template)
}

func (c *CMDProfileCreator) createMainFile() {
	contents := c.getMainFileContents()

	mainFilePath := filepath.Join(c.GetProjectName(), "cmd")
	mainFilePath = filepath.Join(mainFilePath, c.GetProjectName())

	err := os.MkdirAll(mainFilePath, 0755)
	if err != nil {
		c.Log.LogFatalf("Error creating directory '%s' and it's subdirectories", mainFilePath)
	}
	mainFilePath = filepath.Join(mainFilePath, "main.go")

	writeFile(c.Log, mainFilePath, contents)
}

func (c *CMDProfileCreator) getMainFileContents() string {

	return populateTemplate(c.Log, mainFileTemplateStringCMD, c.Template)

}

func (c *CMDProfileCreator) createMakeFile() {
	contents := c.getMakefileContents()

	makeFilePath := filepath.Join(c.GetProjectName(), "Makefile")

	writeFile(c.Log, makeFilePath, contents)

}
func (c *CMDProfileCreator) getMakefileContents() string {

	return populateTemplate(c.Log, makeFileTemplateStringCMD, c.Template)

}
