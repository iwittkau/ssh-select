package putty

import (
	"fmt"
	"os"
	"os/exec"

	sshselect "github.com/5FeetUnder/ssh-select"
)

func NewSSHTerminalWindow(server sshselect.Server) error {
	var cmd *exec.Cmd
	if server.Port == "" {
		cmd = exec.Command("putty", fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	} else {
		cmd = exec.Command("putty", "-P", server.Port, fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
