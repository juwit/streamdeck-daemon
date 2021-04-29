# 0004. Use libxdo instead of xdotool as subcommand

Date: 29 apr. 2021

Status: active

Deprecates: [ADR-0002](0002-use-xdotool-as-subcommand.md)

## Context

As for now, sending keys when button press is made using xdotool :

```go
if button.Write != "" {
    go exec.Command("xdotool", "type", "--delay", "0", button.Write).Start()
}
```

This spawns a new process, and add a command-line dependency.

## Decision

Write a simple wrapper for libxdo, and call it directly

## Consequences

This may improve performance as no new process will be spawned.