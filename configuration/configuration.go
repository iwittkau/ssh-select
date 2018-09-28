package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"

	"github.com/iwittkau/ssh-select"
	"gopkg.in/yaml.v2"
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
		system = "gnome"
	default:
		system = ""
	}

	c := sshselect.Configuration{
		StayOpen: false,
		System:   system,
		Servers: []sshselect.Server{
			sshselect.Server{
				Name:      "Server 1 (Example)",
				IPAddress: "192.168.0.1",
				Username:  "username",
				Profile:   "Homebrew",
			},
			sshselect.Server{
				Name:      "Server 2 (Example)",
				IPAddress: "192.168.0.2",
				Username:  "username",
				Profile:   "Homebrew",
			},
		},
	}
	return WriteToUserHomeDir(&c)
}
