# 0002. Use xdotool as subcommand

Date: 16 mar. 2021

Status: deprecated by [ADR-0004](0004-use-lib-xdo-instead-of-xdotool.md)

## Context

One of the requirements is that a button press could write text to
the currently focused window (to enable basic automation, macros, and emojis input).

I've been experimenting with Robotgo (https://github.com/go-vgo/robotgo).
This works well except that it doesn't support emojis

## Decision

Use directly xdotool to write text

* no need to fight with Robotgo
* emoji support works
* is an external dependency, but is already installed on all my machines

## Consequences

1. call sub-process xdotool to write text
2. external dependencies for other users (that may never exist)