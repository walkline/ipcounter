package main

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/walkline/ipcounter"
	"github.com/walkline/ipcounter/serving/logsinput"
	"github.com/walkline/ipcounter/serving/metrics"
)

func main() {
	logsInputPort := envWithDefault("LOGS_INPUT_PORT", "5000")
	metricsPort := envWithDefault("METRICS_PORT", "9102")

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	index := ipcounter.NewIPv4BucketIndex()

	logsInputServer := logsinput.NewServer(index, log.With(logger, "module", "logs_input"))
	metricsServer := metrics.NewServer(index, log.With(logger, "module", "metrics"))

	go func() {
		logger.Log("logs_input", "Starting server...")
		// TODO: add graceful shutdown
		if err := logsInputServer.ListenAndServe(":" + logsInputPort); err != nil {
			logger.Log("logs_input", err.Error())
		}
	}()

	logger.Log("metrics", "Starting server...")
	// TODO: add graceful shutdown
	if err := metricsServer.ListenAndServe(":" + metricsPort); err != nil {
		logger.Log("metrics", err.Error())
	}
}

func envWithDefault(key string, def string) string {
	result := os.Getenv(key)
	if result == "" {
		return def
	}

	return result
}
