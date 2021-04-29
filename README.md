# streamdeck-daemon

a simple streamdeck tool, written in golang, 
which loads configuration from a JSON file, 
and exposes an HTTP API to dynamically interact with the streamdeck.

it #WorksOnMyMachine.

## requirements

* [x] : write text on button press
* [x] : execute command on button press
* [x] : show images on buttons
* [x] : manage "pages"
* [x] : change page on button press
* [x] : update button on http post

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
* https://github.com/go-chi/chi : for http routing
* libxdo-dev : for sending keyboard keys on button press (sudo apt install libxdo-dev)

```shell
sudo apt install xdotool
```
