package osascript

import (
	"fmt"

	"github.com/iwittkau/ssh-select"

	"github.com/everdev/mack"
)

func NewSSHTerminalWindow(server sshselect.Server) error {
	var err error
	if server.Port == "" {
		_, err = mack.Tell("Terminal", fmt.Sprintf("do script(\"ssh %s@%s\")", server.Username, server.IPAddress))
	} else {
		_, err = mack.Tell("Terminal", fmt.Sprintf("do script(\"ssh -p %s %s@%s\")", server.Port, server.Username, server.IPAddress))
	}

	return err
}

func SetFrontmostTerminalWindowToProfile(profile string) error {
	_, err := mack.Tell("Terminal", fmt.Sprintf("set current settings of front window to (first settings set whose name is \"%s\")", profile))
	return err
}
