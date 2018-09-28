package sshselect

// System Constans
const (
	SystemMacOS = "macos"
	SystemGnome = "gnome"
)

// Configuration ist the struct that holds the SSH-Select configuration located at the users home path
type Configuration struct {
	System   string
	StayOpen bool
	Servers  []Server
}

// Server holds all ssh related information to assemble the ssh command
type Server struct {
	Name      string
	IPAddress string
	Username  string
	Profile   string
	Index     int `yaml:",omitempty"`
}

// Frontend is the needs to be implemented by a cli frontend
type Frontend interface {
	Draw(selected int) (int, bool, error)
}
