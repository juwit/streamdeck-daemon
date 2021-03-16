# streamdeck-daemon

a simple streamdeck tool, written in golang, 
which loads configuration from a JSON file, 
and exposes an HTTP API to dynamically interact with the streamdeck.

## requirements

[ ] : write text on button press
[ ] : execute command on button press
[ ] : show images on buttons
[ ] : manage "pages"
[ ] : change page on button press
[ ] : update button on http post

## building

```shell
go build
```

## running

```shell
./streamdeck-daemon
```

## dependencies

* github.com/magicmonkey/go-streamdeck : for basic streamdeck interaction
* xdotool : for sending keyboard keys on button press

```shell
sudo apt install xdotool
```