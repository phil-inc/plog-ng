# plog-ng

plog-ng is a logging utility.

## Overview

The package consists of the following methods:

### Info

```go
func Info(s string)
```

Info logs a message at level Info. It takes string that is to be logged as an input.

### Infof

```go
func Infof(s string, args ...interface{})
```

Infof logs a message at level Info. It takes string with format specifier and arguments, and logs the resulting string

### Debug

```go
func Debug(s string)
```

Debug logs a message at level Debug. It takes string that is to be logged as an input.

### Debugf

```go
func Debugf(s string, args ...interface{})
```

Debugf logs a message at level Debug. It takes string with format specifier and arguments, and logs the resulting string

### Warn

```go
func Warn(s string)
```

Warn logs a message at level Warn. It takes string that is to be logged as an input.

### Warnf

```go
func Warnf(s string, args ...interface{})
```

Warnf logs a message at level Warn. It takes string with format specifier and arguments, and logs the resulting string

### Error

```go
func Error(s string)
```

Error logs a message at level Error. It takes string that is to be logged as an input.

### Errorf

```go
func (entry *entry) Errorf(s string, args ...interface{})
```

Errorf logs a message at level Error. It takes string with format specifier and arguments, and logs the resulting string

### Panic

```go
func Panic(s string)
```

Panic logs a message at level Panic. It takes string that is to be logged as an input.

### Panicf

```go
func (entry *entry) Panicf(s string, args ...interface{})
```

Panicf logs a message at level Panic. It takes string with format specifier and arguments, and logs the resulting string

## WithFields

```go
func WithFields(fields Fields) *entry
```

WithFields creates an entry from the standard logger and adds multiple fields to it.
Note that it doesn't log until you call Info, Debug, Warn or error on the Entry it returns.

## WithCallerLevel

```go
func WithCallerLevel(level int) *entry
```

WithCallerLevel changes the number of stack frames to ascend for when providing file and line number information about function invocations.
### Example Usage

```go
package main

import (
    logger "github.com/phil-inc/plog-ng/pkg/core"
)

func main() {
    logger.Init() // Logger should always be initialized at the beginning
    logger.Errorf("Error: %s", "New error log")
    logger.WithFields(logger.Fields{"Field": "Additional fields"}).Infof("Info message: %s", "special info message")
}

```
