package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hako/durafmt"
	"github.com/iwittkau/ssh-select"

	"strconv"

	"github.com/iwittkau/ssh-select/configuration"
	"github.com/iwittkau/ssh-select/gnome"
	"github.com/iwittkau/ssh-select/gocui"
	"github.com/iwittkau/ssh-select/metric"
	"github.com/iwittkau/ssh-select/osascript"
	flag "github.com/ogier/pflag"
)

const cpm = int64(200)

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
		metric.InitMetricFile()
		fmt.Println("\nAn example configuration file has been created at '~/.sshs-config'.\nPlease open it an check the 'system' setting. Use 'gnome' if you are on Linux or 'macos' if you use macOS.\n ")
		return
	} else if err == nil && init {
		fmt.Println("  --init ignored because existing configuration would be overwritten.")
		return
	}

	switch config.System {
	case sshselect.SystemMacOS, sshselect.SystemGnome:
		break
	case "":
		fmt.Println("\nSystem not set! Please open '~/.sshs-config' an set the 'system' setting to 'gnome' or 'macos' depending on which of the one you currently use.\n ")
		return
	default:
		fmt.Printf("\n~/.sshs-config: setting 'system: %s' not supported!\nPlease use 'gnome' or 'macos' for 'system'.\n\n", config.System)
		return
	}

	cm, _ := metric.Load()
	if cm == nil {
		metric.InitMetricFile()
	}

	defer func() {
		if cm != nil {
			saved := cm.Count()
			cm.Persist()
			dur := time.Duration(time.Duration(saved/cpm*60) * time.Second)
			if dur >= time.Minute*5 {
				fmt.Printf("ssh-select has saved you approximately %s in total.\n", durafmt.Parse(dur).String())
			}
		}
	}()

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

		for {

			var selected bool

			f, err := gocui.New(config)

			if err != nil {
				fmt.Println("", err)
				return
			}

			i, selected, err = f.Draw(i)

			if err != nil {
				fmt.Println("Frontend:", err.Error())
				return
			} else if !selected {
				return
			}

			switch config.System {

			case sshselect.SystemMacOS:

				err = osascript.NewSSHTerminalWindow(config.Servers[i].Username, config.Servers[i].IPAddress)
				if err != nil {
					fmt.Println("Error:", err)
				}

				err = osascript.SetFrontmostTerminalWindowToProfile(config.Servers[i].Profile)
				if err != nil {
					fmt.Println("Error:", err)
				}

			case sshselect.SystemGnome:
				err = gnome.NewSSHTerminalWindow(config.Servers[i].Username, config.Servers[i].IPAddress, config.Servers[i].Profile)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}

			if cm != nil {
				cm.Add(config.Servers[i].Username+"@"+config.Servers[i].IPAddress, "")
				cm.Persist()
			}

			if !config.StayOpen {
				break
			}

		}
	}

}
