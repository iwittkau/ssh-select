package sshselect

type Configuration struct {
	Servers []Server
}

type Server struct {
	Name      string
	IpAddress string
	Username  string
	Profile   string
	Index     int `yaml:",omitempty"`
}

type Frontend interface {
	Draw() (int, bool, error)
}
