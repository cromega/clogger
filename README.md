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
  logger := clogger.CreateLogger(clogger.Debug)
  logger.AddTarget(clogger.CreateWriterTarget(io.Stdout))
  logger.Debug("my god, it's full of %v", "stars")
}
```
`AddTarget` takes a `*clogger.CloggerTarget`

it produces `10/24/2014 19:45:02, D: my god, it's full of stars`

log levels are: `Debug, Info, Warning, Error, Fatal`
