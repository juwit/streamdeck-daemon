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
* github.com/go-vgo/robotgo : for sending keyboard keys on button press

### robotgo dependencies

robotgo requires a lot of external dependencies to work:

```shell
sudo apt install gcc libc6-dev

sudo apt install libx11-dev xorg-dev libxtst-dev libpng++-dev

sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt install libxkbcommon-dev

sudo apt install xsel xclip
```