package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

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
