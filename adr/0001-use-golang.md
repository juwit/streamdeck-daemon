# 0001. Use Golang as the language

Date: 12 mar. 2021

## Context

Need to find some language that interacts well with an USB device,
probably with some streamdeck library already exising.

I've been experimenting with Java (with lib https://github.com/VVEIRD/StreamDeckCore),
and NodeJS (with lib https://www.npmjs.com/package/elgato-stream-deck)

I also want the daemon to be a simple executable, that I will run as a system service.

## Decision

Choosed Golang because:

* a basic library already exists
* I want to learn something new
* compiles as a simple binary

## Consequences

1. this project starts
2. I have to learn basic golang