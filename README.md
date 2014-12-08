clogger
=======

a very basic logger for Go that doesn't actually terminate the process when you log something Fatal.

## Usage

```go
package main

import (
  "github.com/cromega/clogger"
  "os"
)

func main() {
  var logger = clogger.Logger
  target := os.Getenv("LOG")
  if target == "local" {
    logger = clogger.CreateIoLogger(os.Stdout)
  } else {
    logger = clogger.CreateSyslog("udp", "logs2.papertrailapp:12345", "app")
  }

  logger.Info("logging is awesome")
}
```

This logger does not do any level filtering. That is the concern of the consumer. Levels are Debug, Info, Warning, Error, Fatal (Critical)
