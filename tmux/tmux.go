package tmux

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	sshselect "github.com/5FeetUnder/ssh-select"
)

const (
	defaultSessionName = "sshs"
)

// NewSSHTerminalWindow creates a new tmux window with the ssh command for the given server
func NewSSHTerminalWindow(server sshselect.Server) error {

	var (
		tmuxOut  []byte
		err      error
		session  string
		window   = fmt.Sprintf("%s@%s", server.Username, server.Name)
		detached = false
		sshCmd   string
	)

	if server.Port == "" {
		sshCmd = fmt.Sprintf("ssh %s@%s\n", server.Username, server.IPAddress)
	} else {
		sshCmd = fmt.Sprintf("ssh -p %s %s@%s\n", server.Port, server.Username, server.IPAddress)
	}

	cmd := exec.Command("tmux", "display-message", "-p", "#S")

	if tmuxOut, err = cmd.Output(); err != nil {
		cmd = exec.Command("tmux", "new", "-s", defaultSessionName, "-d", "sshs")
		if err := cmd.Run(); err != nil {
			return err
		}
		session = defaultSessionName
	} else {
		scanner := bufio.NewScanner(strings.NewReader(string(tmuxOut)))
		if scanner.Scan() {
			session = scanner.Text()
		}
	}

	if os.Getenv("TMUX") == "" {
		detached = true
	}

	cmd = exec.Command("tmux", "new-window", "-a", "-t", session, "-n", window)
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("tmux", "send-keys", "-t", session+":"+window, sshCmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	if detached {
		if session != defaultSessionName {
			fmt.Printf("not attached to a tmux session; please attach to your last session with 'tmux a -t %s'\n", session)
		} else {
			fmt.Println("not attached to a tmux session; a new sshs tmux session was created; attach with 'tmux a -t sshs'")
		}
		return sshselect.ErrTmuxNotAttached

	}

	return nil
}
