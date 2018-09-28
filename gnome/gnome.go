package gnome

import (
	"fmt"
	"os/exec"
)

// NewSSHTerminalWindow opens a new gnome-terminal window
// gnome-terminal --profile=[PROFILE_NAME] -- [COMMAND]
func NewSSHTerminalWindow(username, server, profile string) error {
	cmd := exec.Command("gnome-terminal", fmt.Sprintf("--profile=%s", profile), "--", "ssh", fmt.Sprintf("%s@%s", username, server))
	return cmd.Run()
}
