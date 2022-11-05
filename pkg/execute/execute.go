package execute

import (
	"github.com/pspiagicw/groom-create/pkg/log"
	"os/exec"
)

func Execute(log log.Logger, command string, args []string, dir string) (string, error) {

	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	log.LogStep("Executing '%s'", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
