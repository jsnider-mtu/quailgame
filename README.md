# Quail Game

A 2D game based in the Quail Kingdom

## Build commands
### Linux

```
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o quail_game_demo --tags linux
```

### Windows
#### If building on a Linux machine install mingw dependency

```
sudo apt-get install gcc-mingw-w64
```

#### Then run:

```
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o quail_game_demo.exe
```

### Mac OSX

Haven't build for Mac yet
