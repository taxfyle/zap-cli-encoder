package main

import (
	"flag"

	zapclilogger "github.com/taxfyle/zap-cli-logger"
	"go.uber.org/zap"
)

func main() {
	verbose := flag.Bool("verbose", false, "pring verbose output")
	flag.Parse()

	logger, err := zapclilogger.New(*verbose)
	if err != nil {
		panic(err)
	}

	logger = logger.With(zap.String("some", "context"))
	logger.Debug("you should only see this if -verbose is set, and you should see it with context")
	logger.Info("you should always see this, and if -verbose is set you should see it with context")
	logger.Warn("you should always see this with context")
	logger.Error("you should always see this with context")
}
