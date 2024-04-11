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


## Usage Example

This game is designed to be a fun multiplayer experience. Here's how you can get started:

1. **Install the game**: Follow the installation instructions above to get the game up and running on your system.

2. **Start the game**: Once installed, you can start the game by running the following command in your terminal:

    ```shell
    make run_game
    ```

3. **Connect to a server**: The game is multiplayer, so you'll need to connect to a server. If you're running your own server, you can connect to it using the following command:

    ```shell
    todo
    ```

   Replace `localhost:8080` with the address of your server.

4. **Play the game**: Now that you're connected, you can start playing the game with your friends!

5. **Customize your experience**: The game is fully customizable and open source. You can add your own features, and tweak the game rules. To do this, you'll need to dive into the code. Check out the `game` directory for the game logic, and the `assets` directory for the game's graphics and sound effects.

Remember, this is a community project, so feel free to contribute! If you make a cool feature or level, consider making a pull request so everyone can enjoy your work.

## Release History

* 0.1.0
    * Base multiplayer version 
* 0.0.1
    * Work in progress


## Contributing

1. Fork it (<https://github.com/zumosik/online_game/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

