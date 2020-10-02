# gologger
[![Build Status](https://travis-ci.org/mldev-ai/gologger.svg?branch=main)](https://travis-ci.org/mldev-ai/gologger)
[![Go Report Card](https://goreportcard.com/badge/github.com/mldev-ai/gologger)](https://goreportcard.com/report/github.com/mldev-ai/gologger)

Gologger is simple scalable logger for any type of application logging. Ability to have different log levels with corresponding environment.

```shell script
# To install in your project
go get github.com/mldev-ai/gologger
```

### Easy to Use
```go
package main

import "github.com/mldev-ai/gologger"

const LOG_LEVEL=3

func main() {
    
    myLogger := gologger.NewGoLogger(INFO_LEVEL, false, "YOUR_SCOPE")
    myLogger.Info("This is some log!!!")
    myLogger.Info(1)
    
}
```
