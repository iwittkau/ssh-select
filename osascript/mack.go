package osascript

import (
	"fmt"

	"github.com/everdev/mack"
	"github.com/iwittkau/ssh-select"
)

// NewSSHTerminalWindow opens a new macos Terminal.app window for the selected server
func NewSSHTerminalWindow(server sshselect.Server) error {
	var err error
	if server.Port == "" {
		_, err = mack.Tell("Terminal", fmt.Sprintf("do script(\"ssh %s@%s\")", server.Username, server.IPAddress))
	} else {
		_, err = mack.Tell("Terminal", fmt.Sprintf("do script(\"ssh -p %s %s@%s\")", server.Port, server.Username, server.IPAddress))
	}

	return err
}

// SetFrontmostTerminalWindowToProfile sets the active window to desired profile
func SetFrontmostTerminalWindowToProfile(profile string) error {
	_, err := mack.Tell("Terminal", fmt.Sprintf("set current settings of front window to (first settings set whose name is \"%s\")", profile))
	return err
}

// NewSSHITermWindow tells iTerm2 to create a new window with the ssh-command and desired profile
func NewSSHITermWindow(server sshselect.Server, profile string) error {
	var err error
	if server.Port == "" {
		_, err = mack.Tell("iTerm2", fmt.Sprintf("create window with profile \"%s\" command \"ssh %s@%s\"", profile, server.Username, server.IPAddress))
	} else {
		_, err = mack.Tell("iTerm2", fmt.Sprintf("create window with profile \"%s\" command \"ssh -p %s %s@%s\"", profile, server.Port, server.Username, server.IPAddress))
	}
	return err
}

// NewSSHITermTab tells iTerm2 to create a new tab with the ssh-command and desired profile
func NewSSHITermTab(server sshselect.Server, profile string) error {
	var err error
	if server.Port == "" {
		_, err = mack.Tell("iTerm2", fmt.Sprintf("tell current window\ncreate tab with profile \"%s\" command \"ssh %s@%s\"\nend tell", profile, server.Username, server.IPAddress))
	} else {
		_, err = mack.Tell("iTerm2", fmt.Sprintf("tell current window\ncreate tab with profile \"%s\" command \"ssh -p %s %s@%s\"\nend tell", profile, server.Port, server.Username, server.IPAddress))
	}
	return err
}
