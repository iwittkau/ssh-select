package osascript

import (
	"fmt"
	"os/exec"

	"github.com/5FeetUnder/ssh-select/configuration"

	sshselect "github.com/5FeetUnder/ssh-select"
)

// NewSSHTerminalWindow opens a new macos Terminal.app window for the selected server
func NewSSHTerminalWindow(server sshselect.Server) error {
	sshCmd := configuration.BuildSSHCmdTunnelConfigStr(server.Tunnel)
	var err error
	if server.Port == "" {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s %s@%s\"", sshCmd, server.Username, server.IPAddress)).Run()
	} else {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s -p %s %s@%s\"", sshCmd, server.Port, server.Username, server.IPAddress)).Run()
	}
	return err
}

// SetFrontmostTerminalWindowToProfile sets the active window to desired profile
func SetFrontmostTerminalWindowToProfile(profile string) error {
	err := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Terminal\" to set current settings of front window to (first settings set whose name is \"%s\")", profile)).Run()
	return err
}

// NewSSHITermWindow tells iTerm2 to create a new window with the ssh-command and desired profile
func NewSSHITermWindow(server sshselect.Server, profile string) error {
	var err error
	sshCmd := configuration.BuildSSHCmdTunnelConfigStr(server.Tunnel)
	if server.Port == "" {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"iTerm2\" to create window with profile \"%s\" command \"%s %s@%s\"", profile, sshCmd, server.Username, server.IPAddress)).Run()
	} else {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"iTerm2\" to create window with profile \"%s\" command \"%s -p %s %s@%s\"", profile, sshCmd, server.Port, server.Username, server.IPAddress)).Run()
	}
	return err
}

// NewSSHITermTab tells iTerm2 to create a new tab with the ssh-command and desired profile
func NewSSHITermTab(server sshselect.Server, profile string) error {
	var err error
	sshCmd := configuration.BuildSSHCmdTunnelConfigStr(server.Tunnel)
	if server.Port == "" {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"iTerm2\" to tell current window to create tab with profile \"%s\" command \"%s %s@%s\"", profile, sshCmd, server.Username, server.IPAddress)).Run()
	} else {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"iTerm2\" to tell current window to create tab with profile \"%s\" command \"%s -p %s %s@%s\"", profile, sshCmd, server.Port, server.Username, server.IPAddress)).Run()
	}
	return err
}
