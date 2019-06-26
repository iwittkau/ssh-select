package gnome

import (
	"fmt"
	"os/exec"

	sshselect "github.com/5FeetUnder/ssh-select"
)

// NewSSHTerminalWindow opens a new gnome-terminal window
// gnome-terminal --profile=[PROFILE_NAME] -- [COMMAND]
func NewSSHTerminalWindow(server sshselect.Server) error {
	var cmd *exec.Cmd
	if server.Port == "" {
		cmd = exec.Command("gnome-terminal", fmt.Sprintf("--profile=%s", server.Profile), "--", "ssh", fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	} else {
		cmd = exec.Command("gnome-terminal", fmt.Sprintf("--profile=%s", server.Profile), "--", "ssh", "-p", server.Port, fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	}
	return cmd.Run()
}
