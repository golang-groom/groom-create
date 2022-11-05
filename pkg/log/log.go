package log

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

// Custom Logger interface
// Logs with colors!
// Usage same as standard library logger
type Logger interface {
	LogFatalf(string, ...interface{})
	LogSuccess(string, ...interface{})
	LogStep(string, ...interface{})
	LogInfo(string, ...interface{})
}

// Instance of the Logger interface
type DefaultLogger struct {
}

// Log when something goes wrong
func (d *DefaultLogger) LogFatalf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	color.Red(message)
	os.Exit(1)

}

// Log when something succeeds
func (d *DefaultLogger) LogSuccess(format string, v ...interface{}) {
	message := "[*] ✓ " + fmt.Sprintf(format, v...)
	color.Green(message)
}

// Log a single step
func (d *DefaultLogger) LogStep(format string, v ...interface{}) {
	message := "[*] + " + fmt.Sprintf(format, v...)
	color.Yellow(message)
}

// Log basic things. Only used when `-verbose` option used
func (d *DefaultLogger) LogInfo(format string, v ...interface{}) {}
