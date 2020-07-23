run:
	@go run cmd/ipcounter/main.go

test:
	@go test $(shell go list ./...)

bench:
	@go test $(shell go list ./...) -bench=.

loadtest:
	@./vegeta-targets-generator.sh | vegeta attack -format=json -rate=0 -max-workers=100 -duration=30s | vegeta report
