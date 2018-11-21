package main

import (
	"go.uber.org/zap"
)

func main() {
	url := "http://www.baidu.com"
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
