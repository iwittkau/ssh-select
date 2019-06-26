package sshselect

// System Constans
const (
	SystemMacOS = "macos"
	SystemGnome = "gnome"
	SystemITerm = "iterm"
	SystemTmux  = "tmux"
	SystemPlink = "plink"
)

// Configuration ist the struct that holds the SSH-Select configuration located at the users home path
type Configuration struct {
	System   string
	StayOpen bool
	UseTabs  bool `yaml:",omitempty"`
	Servers  []Server
}

// Server holds all ssh related information to assemble the ssh command
type Server struct {
	Name      string
	IPAddress string
	Username  string
	Profile   string
	Port      string                `yaml:",omitempty"`
	Index     int                   `yaml:",omitempty"`
	Tunnel    []TunnelConfiguration `yaml:",omitempty"`
}

// TunnelConf describes a tunneling configuration
type TunnelConfiguration struct {
	Port     string
	HostPort string
	Host     string
}

// Frontend is the needs to be implemented by a cli frontend
type Frontend interface {
	Draw(selected int) (int, bool, error)
}
