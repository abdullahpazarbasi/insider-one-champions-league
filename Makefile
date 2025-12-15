.DEFAULT_GOAL := setup
SHELL := /bin/bash

.PHONY: setup stopped status-checked cleaned-up service-logs-viewed service-shell-connected service-test service-coverage bff-logs-viewed bff-shell-connected bff-test bff-coverage spa-logs-viewed spa-shell-connected spa-dev spa-test spa-coverage

setup:
	@bash scripts/setup.sh

stopped:
	@bash scripts/stop.sh

status-checked:
	@bash scripts/check-status.sh

service-logs-viewed:
	@bash scripts/view-service-log.sh

service-shell-connected:
	@bash scripts/execute-service-shell.sh

service-test:
	@cd back/service && go test -v ./...

service-coverage:
	@cd back/service && go test -coverprofile=coverage.out ./...

bff-logs-viewed:
	@bash scripts/view-bff-log.sh

bff-shell-connected:
	@bash scripts/execute-bff-shell.sh

bff-test:
	@cd back/bff && go test -v ./...

bff-coverage:
	@cd back/bff && go test -coverprofile=coverage.out ./...

spa-logs-viewed:
	@bash scripts/view-spa-log.sh

spa-shell-connected:
	@bash scripts/execute-spa-shell.sh

spa-dev:
	@cd front/spa && npm run dev

spa-test:
	@cd front/spa && npm run test

spa-coverage:
	@cd front/spa && npm run test:coverage

cleaned-up:
	@echo "🔥 This will terminate everything 🔥 Continue? (y/N)"
	@read -r confirm && [[ $$confirm == "y" ]] && bash scripts/clean-up.sh || echo "Cleaning up cancelled"
