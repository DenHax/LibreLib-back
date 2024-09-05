# Backend for LibreLib.

## API:

## Stack:
Golang 1.21.9


## AUTOSTART with one command (for developers and viewers):

Dependencies: git, make, docker, shell (bash or powershell)

### Usage:
```bash
git clone https://github.com/DenHax/LibreLib-back

cd LibreLib-back

make auto-start
```

What's happend?

After enter `make auto-start` script choose target OS, choose follow script for autostart from scripts/ and after run option `compose` for target system.

Autostart generate and activate dotenv file for correct operation of the application - add required environment variables to local machine for example how it start.  
Compose activate deployments/compose.yaml and build all services: backend app, db server.

 <!--TODO: nginx, systemd unit-->

### Windows

1. Docker via official site

Donwload and install Docker Desktop from https://www.docker.com/products/docker-desktop

2. Make and Git via Git Bash

Install Git for Windows from https://gitforwindows.org/

Make will be avainble by default. Create alias in your powershell config on "make"

3. Docker, Make, Git via Chocolatey
```bash
choco install make git docker-desktop
```

4. Make and Git Scoop
```bash
scoop install make git
```

5. Use WSL (Windows Subsystem for Linux)

Install wsl with latest Ubuntu (or choose your preferred distribution)

```bash
wsl --install
```

 Update system and install make with your distribution's package manager
```bash
sudo apt update
sudo apt install make git docker.io

```
6. Make via Cygwin

Download and install Cygwin from https://www.cygwin.com/ and choose make in installing proccess

7. Make via MinGW

Download and install MinGW from https://osdn.net/projects/mingw/releases/ and choose make in installing proccess

### Debian/Ubuntu, Mint, Pop! OS, Zorin
```bash
sudo apt install make git docker
```

### Fedora, CentOS
```bash
sudo dnf install make git docker
```
### Astra Linux

```bash
sudo apt install make git docker
```

### ALT Linux

```bash
sudo apt-get install make git docker
```

# Install make in Arch, Manjaro
```bash
sudo pacman -S make git docker
```

### Install make in VoidLinux

```bash
sudo xbps-install -S make git docker
```

### NixOS and Nix

# 1. Install via nix-env
```bash
nix-env -iA nixpkgs.make nixpkgs.git nixpkgs.docker
```

### 2. Install in system packages

```nix
# /etc/nixos/configuration.nix
{ 
    ...
    environment.systemPackages = with pkgs; [ make git docker ];
    ...
}
```

Rebuild system

```bash
sudo nixos-rebuild switch
```

### Alpine

```bash
sudo apk add make git docker
```

### OpenSUSE

```bash
sudo zypper install make git docker
```

### Slackware

```bash
sudo slackpkg install make git docker
```

### FreeBSD

```bash
pkg install gmake git docker
```

### Deepin

```bash
sudo apt install make git docker
```

## Instalation + usage:


## Manual


```bash
### Install golang

apt install go

### Install dependencies

git clone https://github.com/DenHax/LibreLib-back
cd LibreLib-back
go mod tidy

# Run server
go run ./cmd/librelib/main.go

# Build app
go build ./cmd/librelib/main.go 
```

### Docker (compose):
```bash
apt install docker

git clone https://github.com/DenHax/LibreLib-back

cd LibreLib-back

docker compose -f ./deployments/compose.yaml up --build
```
