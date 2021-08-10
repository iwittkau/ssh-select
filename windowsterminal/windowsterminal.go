package windowsterminal

import (
	"fmt"
	"os/exec"

	sshselect "github.com/iwittkau/ssh-select"
)

// NewSSHTerminalWindow opens a new tab in Windows Terminal running SSH for the selected server
func NewSSHTerminalWindow(server sshselect.Server) error {
	var cmd *exec.Cmd
	if server.Port == "" {
		cmd = exec.Command("wt", "-p", server.Profile, "ssh", fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	} else {
		cmd = exec.Command("wt", "-p", server.Profile, "ssh", "-P", server.Port, fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	}

	return cmd.Run()
}

// NewSSHTerminalWindowFix does the same as NewSSHTerminalWindow, but uses cmd to invoke wt. Used for when wt.exe cannot be found in PATH (apparently exec.LookPath doesnt find files that are in AppData)
func NewSSHTerminalWindowFix(server sshselect.Server) error {
	var cmd *exec.Cmd
	if server.Port == "" {
		cmd = exec.Command("cmd", "/C", "wt", "-p", server.Profile, "ssh", fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	} else {
		cmd = exec.Command("cmd", "/C", "wt", "-p", server.Profile, "ssh", "-P", server.Port, fmt.Sprintf("%s@%s", server.Username, server.IPAddress))
	}

	return cmd.Run()
}