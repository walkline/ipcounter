package main

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/walkline/ipcounter"
	"github.com/walkline/ipcounter/serving/logsinput"
)

func main() {
	logsInputPort := envWithDefault("LOGS_INPUT_PORT", "5000")

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	index := ipcounter.NewIPv4BucketIndex()

	logsInputServer := logsinput.NewServer(index, log.With(logger, "module", "logs_input"))

	logger.Log("logs_input", "Strating server...")

	// TODO: add graceful shutdown
	if err := logsInputServer.ListenAndServe(":" + logsInputPort); err != nil {
		logger.Log("err", "logs input ListenAndServe err: "+err.Error())
	}
}

func envWithDefault(key string, def string) string {
	result := os.Getenv(key)
	if result == "" {
		return def
	}

	return result
}
