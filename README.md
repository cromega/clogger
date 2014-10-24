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
  logger := clogger.InitLogger(os.Stdout, clogger.Debug) // it takes any io.Writer
  logger.Debug("my god, it's full of %v", "stars")
})
```

it produces `10/24/2014 19:45:02, D: my god, it's full of stars`

log levels are: `Debug, Info, Warning, Error, Fatal`
