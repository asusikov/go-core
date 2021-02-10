PKG := goondex
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

run:
	@go run goondex/cli

test:
	@go test -short ${PKG_LIST}
