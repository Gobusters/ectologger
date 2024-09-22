# ectologger

ectologger is a Go library that provides a standardized logging interface, allowing you to decouple your application code from specific logging implementations.

## Features

- Standardized logging interface
- Multiple log levels (Debug, Info, Warn, Error, Fatal)
- Context-aware logging
- Structured logging with fields
- Easy integration with existing loggers (e.g., zap)
- Customizable log output format

## Installation

Add ectologger to your project:

```bash
go get -u github.com/Gobusters/ectologger
```

## Usage

1. Create a logger instance:

   ```go
   logger := ectologger.NewDefaultEctoLogger()
   ```

2. Use the logger in your code:

   ```go
   logger.Info("Starting the application")
   logger.WithField("request_id", "12345").Info("Handling request")
   ```

## Zap adapter

ectologger can be easily integrated with existing logging libraries like zap:

1. Install zap:

```bash
go get -u go.uber.org/zap
```

2. Use the zap adapter:

```go

import (
	"github.com/Gobusters/ectologger/zapadapter"
	"go.uber.org/zap"
)


zapLogger, _ := zap.NewProduction()
ectoLogger := zapadapter.NewZapEctoLogger(zapLogger, nil)
```

## License

ectologger is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
