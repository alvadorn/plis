package translation

import (
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

func ShellOut(binName string, args []string) string {
	var (
		output []byte
		error  error
	)
	output, error = exec.Command(binName, args...).Output()

	if error != nil {
		color.Red("Command:", binName+" "+strings.Join(args, " "))
		color.Red("Error:", error)
		os.Exit(1)
	}

	return string(output)
}
