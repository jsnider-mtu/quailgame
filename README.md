# Quail Game

A 2D game based in the Quail Kingdom

## Build commands
### Linux

```
CGO\_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o quail\_game\_demo --tags linux
```

### Windows
#### If building on a Linux machine install mingw dependency

```
sudo apt-get install gcc-mingw-w64
```

#### Then run:

```
CGO\_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86\_64-w64-mingw32-gcc go build -o quail\_game\_demo.exe
```

### Mac OSX

Haven't build for Mac yet
