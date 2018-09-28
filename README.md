# ssh-select

![sshs demo](.github/sshs.gif)

## Installation

```bash
go get -u github.com/iwittkau/ssh-select
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
```

#### General Options

`system` - name of the system you use, either `macos` for macOS or `gnome` for GNOME terminals are supported.   
`stayopen` - leave `sshs` open after a server selection: `true` or `false`

#### Connection Options

`name` - name of the connection  
`ipaddress` - the IP address to connect to  
`username` -  the username  
`profile` - name of the Terminal.app profile


## Usage

```bash
sshs
```

