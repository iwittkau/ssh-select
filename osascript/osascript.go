package osascript

import (
	"fmt"
	"os/exec"

	"github.com/iwittkau/ssh-select"
)

// NewSSHTerminalWindow opens a new macos Terminal.app window for the selected server
func NewSSHTerminalWindow(server sshselect.Server) error {
	var err error
	if server.Port == "" {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"ssh %s@%s\"", server.Username, server.IPAddress)).Run()
	} else {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"ssh -p %s %s@%s\"", server.Port, server.Username, server.IPAddress)).Run()
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
	if server.Port == "" {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"iTerm2\" to create window with profile \"%s\" command \"ssh %s@%s\"", profile, server.Username, server.IPAddress)).Run()
	} else {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"iTerm2\" to create window with profile \"%s\" command \"ssh -p %s %s@%s\"", profile, server.Port, server.Username, server.IPAddress)).Run()
	}
	return err
}

// NewSSHITermTab tells iTerm2 to create a new tab with the ssh-command and desired profile
func NewSSHITermTab(server sshselect.Server, profile string) error {
	var err error
	if server.Port == "" {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"iTerm2\" to tell current window to create tab with profile \"%s\" command \"ssh %s@%s\"", profile, server.Username, server.IPAddress)).Run()
	} else {
		err = exec.Command("osascript", "-e", fmt.Sprintf("tell application \"iTerm2\" to tell current window to create tab with profile \"%s\" command \"ssh -p %s %s@%s\"", profile, server.Port, server.Username, server.IPAddress)).Run()
	}
	return err
}
