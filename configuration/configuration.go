package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"

	sshselect "github.com/iwittkau/ssh-select"
	yaml "gopkg.in/yaml.v2"
)

// ReadFromUserHomeDir reads the SSH-Select configuration from the user's home directory
func ReadFromUserHomeDir() (*sshselect.Configuration, error) {

	config := sshselect.Configuration{}

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(fmt.Sprintf("%s%s.sshs-config", usr.HomeDir, string(os.PathSeparator)))
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	servers := &config.Servers

	for k, v := range *servers {
		v.Index = k + 1
		(*servers)[k] = v
	}

	return &config, nil

}

// ReadFromWorkingDirectory reads the SSH-Select configuration from the current working directory
func ReadFromWorkingDirectory() (*sshselect.Configuration, error) {

	config := sshselect.Configuration{}

	if _, err := os.Stat(".sshs"); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(".sshs")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	servers := &config.Servers

	for k, v := range *servers {
		v.Index = k + 1
		(*servers)[k] = v
	}

	return &config, nil

}

// WriteToUserHomeDir writes the SSH-Select configuration to the user's home directory
func WriteToUserHomeDir(conf *sshselect.Configuration) error {

	usr, err := user.Current()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s%s.sshs-config", usr.HomeDir, string(os.PathSeparator)), data, os.ModePerm)

	return err
}

// Init initializes a configuration file in the users home directory
func Init() error {

	var system string

	switch runtime.GOOS {
	case "darwin":
		system = "macos"
	case "linux":
		system = "tmux"
	default:
		system = "tmux"
	}

	c := sshselect.Configuration{
		StayOpen: true,
		System:   system,
		Servers: []sshselect.Server{
			sshselect.Server{
				Name:      "Server 1 (Example)",
				IPAddress: "192.168.0.1",
				Username:  "username",
				Profile:   "Default",
			},
			sshselect.Server{
				Name:      "Server 2 (Example)",
				IPAddress: "192.168.0.2",
				Username:  "username",
				Profile:   "Default",
				Tunnel: []sshselect.TunnelConfiguration{
					{
						Port:     "8080",
						HostPort: "80",
						Host:     "localhost",
					},
				},
			},
		},
	}
	return WriteToUserHomeDir(&c)
}

// BuildSSHCmdTunnelConfigStr builds a string with the tunnel configuration attached
func BuildSSHCmdTunnelConfigStr(tun []sshselect.TunnelConfiguration) (tunStr string) {
	tunStr = "ssh"
	if tun != nil {
		if len(tun) > 0 {
			for i := range tun {
				tunStr = fmt.Sprintf("%s -L %s:%s:%s", tunStr, tun[i].Port, tun[i].Host, tun[i].HostPort)
			}
		}
	}
	return
}
