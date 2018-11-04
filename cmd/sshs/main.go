package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hako/durafmt"
	"github.com/iwittkau/ssh-select"
	"github.com/iwittkau/ssh-select/configuration"
	"github.com/iwittkau/ssh-select/gnome"
	"github.com/iwittkau/ssh-select/gocui"
	"github.com/iwittkau/ssh-select/metric"
	"github.com/iwittkau/ssh-select/osascript"
	"github.com/iwittkau/ssh-select/tmux"
	flag "github.com/ogier/pflag"
)

var (
	help = `
usage:  sshs [id]   (id corresponds to an id shown in the terminal ui
                    and can be used for quickstart)

             --init  (creates an example configuration file in the 
                     user's home directory: '~/.ssh-config')

             --version (prints the current version)

        sample configuration (~/.sshs-config):
		
	---
	system: tmux
	stayopen: true
	usetabs: true # currently only used by iTerm2 'system' setting
	servers:
	- name: raspberry
		ipaddress: 192.168.1.2
		username: pi
		profile: Default
		port: 22
	---
		
	supported system: 
		macos (standard macOS Terminal.app)
		gnome (linux running GNOME)
		iterm (iTerm2 on macOS)
		tmux  (tmux, system independent)

	Notice: macOS is a trademark of Apple Inc., registered in the U.S. and other countries.
`

	version = "dev"
	commit  = "none"
	date    = "unknown"
)

const cpm = int64(200)

func main() {
	var init, ver bool
	flag.BoolVar(&init, "init", false, "Creates an example configuration file in the user's home directory: '~/.ssh-config'")
	flag.BoolVar(&ver, "version", false, "Prints the current version")
	flag.Usage = func() {
		fmt.Println(help)
	}
	flag.Parse()

	if ver {
		fmt.Printf("version: %v, commit: %v, built at: %v\n", version, commit, date)
		os.Exit(0)
	}

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
		fmt.Println("\nAn example configuration file has been created at '~/.sshs-config'.\nPlease open it an check the 'system' setting. Refer to 'sshs -h' for supported systems.\n ")
		return
	} else if err == nil && init {
		fmt.Println("  --init ignored because existing configuration would be overwritten.")
		return
	}

	switch config.System {
	case sshselect.SystemMacOS, sshselect.SystemGnome, sshselect.SystemITerm, sshselect.SystemTmux:
		break
	case "":
		fmt.Println("\nSystem not set! Please open '~/.sshs-config' an set the 'system' setting. Refer to 'sshs -h' for supported systems.\n ")
		return
	default:
		fmt.Printf("\n~/.sshs-config: setting 'system: %s' not supported!\nRefer to 'sshs -h' for supported systems.\n\n", config.System)
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

	var (
		i           int
		preselected bool
	)

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

		preselected = true
		i--

	}

	for {

		if !preselected {

			var selected bool

			f, err := gocui.New(config)

			if err != nil {
				fmt.Println("", err.Error())
				return
			}

			i, selected, err = f.Draw(i)

			if err != nil {
				fmt.Println("Frontend:", err.Error())
				return
			} else if !selected {
				return
			}
		}

		switch config.System {

		case sshselect.SystemMacOS:
			err = osascript.NewSSHTerminalWindow(config.Servers[i])
			if err != nil {
				fmt.Println("Error:", err.Error())
			}

			err = osascript.SetFrontmostTerminalWindowToProfile(config.Servers[i].Profile)
			if err != nil {
				fmt.Println("Error:", err.Error())
			}

		case sshselect.SystemGnome:
			err = gnome.NewSSHTerminalWindow(config.Servers[i])
			if err != nil {
				fmt.Println("Error:", err.Error())
			}
		case sshselect.SystemITerm:
			if config.UseTabs {
				err = osascript.NewSSHITermTab(config.Servers[i], config.Servers[i].Profile)
				if err != nil {
					fmt.Println("Error:", err.Error())
				}
			} else {
				err = osascript.NewSSHITermWindow(config.Servers[i], config.Servers[i].Profile)
				if err != nil {
					fmt.Println("Error:", err.Error())
				}
			}
		case sshselect.SystemTmux:
			err = tmux.NewSSHTerminalWindow(config.Servers[i])
			if err == sshselect.ErrTmuxNotAttached {
				os.Exit(0)
			} else if err != nil {
				fmt.Println("Error:", err.Error())
				os.Exit(0)
			}
		}

		if cm != nil {
			if config.Servers[i].Port == "" {
				cm.Add(config.Servers[i].Username+"@"+config.Servers[i].IPAddress, "")
			} else {
				cm.Add("-p "+config.Servers[i].Port+" "+config.Servers[i].Username+"@"+config.Servers[i].IPAddress, "")
			}
			cm.Persist()
		}

		if !config.StayOpen || preselected {
			break
		}

	}

}
