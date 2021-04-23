# 0003. Use xdotool as subcommand

Date: 23 apr. 2021

## Context

For the HTTP API part, I need a mux/router capable of handling multiple requests.
I should be lightweight and easy to use 
(as I'm a noob with Golang, I already fight with the language, don't want to fight with a fwk).

Standard Goland ServerMux seems really limited (basic URLs params are not supported)

Alternatives are : 
* https://github.com/go-chi/chi
* https://github.com/gorilla/mux

Choosed chi out of random

## Decision

Use Chi as the router/middleware.
May switch to Gorilla at some time if I encounter some difficulties.

## Consequences

Chi is now a dependency