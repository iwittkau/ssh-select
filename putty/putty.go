package putty

import (
	"fmt"
	"os/exec"

	sshselect "github.com/iwittkau/ssh-select"
)

// NewSSHTerminalWindow opens a new PuTTY SSH terminal window for the selected server
func NewSSHTerminalWindow(server sshselect.Server) error {
	var cmd *exec.Cmd
	if server.Port == "" {
		cmd = exec.Command("putty", fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	} else {
		cmd = exec.Command("putty", "-P", server.Port, fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	}

	return cmd.Run()
}
