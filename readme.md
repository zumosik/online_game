# Online game
> Multiplayer game from scratch, written fully in Go

[![Built with Go](https://img.shields.io/badge/Built%20with-Go-00ADD8.svg)](https://golang.org/)
![GitHub License](https://img.shields.io/github/license/zumosik/online_game)

A little multiplayer game using TCP protocol.  
Game built with [raylib bindings for go](https://github.com/gen2brain/raylib-go).  
All communication between is done using binary packets.  
For serialization and deserialization is used [bb-marshalling](https://github.com/zumosik/bb-marshaling)
 
![](header.png)

## Installation

First you need to install [Go](https://go.dev/dl/)

### Ubuntu
```shell
apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libwayland-dev libxkbcommon-dev
```

### Fedora
```shell
dnf install mesa-libGL-devel libXi-devel libXcursor-devel libXrandr-devel libXinerama-devel wayland-devel libxkbcommon-devel
```

### MacOS
On MacOS you need Xcode or Command Line Tools for Xcode.

### Windows:
coming soon

## Usage example

You can play game with your friends or add some funny features to existing game

## Release History

* 0.0.1
    * Work in progress


## Contributing

1. Fork it (<https://github.com/zumosik/online_game/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

