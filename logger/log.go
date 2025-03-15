// This package is used to configure the logger for the application.
// The Configure function sets the prefix, flags, and output of the logger.
// The prefix is set to the application name, the flags are set to include the date, time, and short file name, and the output is set to the standard logger output.
package logger

import (
	"fmt"
	"log"
)

// Configure configures the logger for the application.
// It sets the prefix, flags, and output of the logger.
// The prefix is set to the application name, the flags are set to include the date, time, and short file name, and the output is set to the standard logger output.
func Configure(appName string) *log.Logger {
	logger := &log.Logger{}

	logger.SetPrefix(fmt.Sprintf("[%s] ", appName))
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetOutput(log.Writer())

	return logger
}
