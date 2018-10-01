# ssh-select

![sshs demo](.github/sshs.gif)

## Installation

```bash
go get -u github.com/iwittkau/ssh-select
cd $GOPATH/src/github.com/iwittkau/ssh-select/cmd/sshs
go get ./...
go install github.com/iwittkau/ssh-select/cmd/sshs
```

## Setup

```bash
sshs --init
```

Edit your configuration file `~/.sshs-config`.

### Configuration Example

```yml
system: macos
stayopen: true
servers:
- name: raspi-3
  ipaddress: 10.0.0.4
  username: pi
  profile: Homebrew
- name: NAS
  ipaddress: nas.local
  username: nas-admin
  profile: Pro
  port: 2222
```

#### General Options

`system` - name of the system you use, either `macos` for macOS or `gnome` for GNOME terminals are supported.   
`stayopen` - leave `sshs` open after a server selection: `true` or `false`

#### Connection Options

`name` - name of the connection  
`ipaddress` - the IP address to connect to  
`username` -  the username  
`profile` - name of the Terminal.app profile
`port` - set a non-default port


## Usage

```bash
sshs [index]
```

`index` optional number, directly sets up a connection with the corresponding index of your configuration file (also shown in the "ui").

