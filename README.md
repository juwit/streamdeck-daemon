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
go mod tidy
go build
```

## running

```shell
./streamdeck-daemon
```

## installing

```shell
sudo cp ./streamdeck-daemon /opt
sed -i "s/username/$USER/" streamdeck-daemon.service
sudo cp ./streamdeck-daemon.service /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl enable streamdeck-daemon
sudo systemctl start streamdeck-daemon
```

## dependencies

* github.com/magicmonkey/go-streamdeck : for basic streamdeck interaction
* https://github.com/go-chi/chi : for http routing
* libxdo-dev : for sending keyboard keys on button press (sudo apt install libxdo-dev)
