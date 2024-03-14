package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteCommand(command string) (string, error) {

	args := strings.Fields(command)           // Spliting the command string into individual arguments
	cmd := exec.Command(args[0], args[1:]...) // Create the command object

	output, err := cmd.CombinedOutput() // Execute the command
	if err != nil {
		return "", fmt.Errorf("unable to execute command: %v", err)
	}
	return string(output), nil
}
