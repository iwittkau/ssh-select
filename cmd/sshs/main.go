package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/iwittkau/ssh-select/configuration"
	"github.com/iwittkau/ssh-select/gocui"
	"github.com/iwittkau/ssh-select/osascript"
	flag "github.com/ogier/pflag"
)

func main() {
	var init bool
	flag.BoolVar(&init, "init", false, "Creates an example configuration file in the user's home directory: '~/.ssh-config'")

	flag.Parse()

	config, err := configuration.ReadFromUserHomeDir()
	if err != nil && !init {
		fmt.Println("Configuration error: " + err.Error() + ".\nRun 'sshs --help' for more information.")
		return
	} else if err != nil && init {
		if err := configuration.Init(); err != nil {
			fmt.Println("Error creating configuration: " + err.Error())
			return
		}
		fmt.Println("An example configuration file has been created at '~/.sshs-config'.")
		return
	} else if err == nil && init {
		fmt.Println("  --init ignored because existing configuration would be overwritten.")
		return
	}

	var i int

	if len(os.Args) == 2 {

		i, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Not a number: ", os.Args[1])
			return
		}

		if i > len(config.Servers) || i < 1 {
			for _, s := range config.Servers {
				fmt.Println(s.Index, ":", s.Name)
			}
			fmt.Println("Not configured:", i)
			return
		}

		i--

	} else {

		var selected bool

		//f := promptui.New(config)

		f, err := gocui.New(config)

		if err != nil {
			fmt.Println("", err)
			return
		}

		i, selected, err = f.Draw()

		if err != nil {
			fmt.Println("Frontend:", err.Error())
			return
		} else if !selected {
			return
		}
	}

	err = osascript.NewSSHTerminalWindow(config.Servers[i].Username, config.Servers[i].IpAddress)
	if err != nil {
		panic(err)
	}

	err = osascript.SetFrontmostTerminalWindowToProfile(config.Servers[i].Profile)
	if err != nil {
		panic(err)
	}

}
