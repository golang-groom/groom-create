package execute

import (
	"os/exec"

	"github.com/pspiagicw/colorlog"
)

func Execute(log colorlog.ColorLogger, command string, args []string, dir string) (string, error) {

	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	log.LogInfo("Executing '%s'", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
