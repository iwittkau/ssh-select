package osascript

import (
	"fmt"

	"github.com/everdev/mack"
)

func NewSSHTerminalWindow(username, server string) error {
	_, err := mack.Tell("Terminal", fmt.Sprintf("do script(\"ssh %s@%s\")", username, server))
	return err
}

func SetFrontmostTerminalWindowToProfile(profile string) error {
	_, err := mack.Tell("Terminal", fmt.Sprintf("set current settings of front window to (first settings set whose name is \"%s\")", profile))
	return err
}
