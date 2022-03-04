
GO ?= go
GOTESTSUM ?= gotestsum

.PHONY: build
## Monta o plugin
build:
	@$(GO) mod tidy
	@$(GO) build -buildmode=plugin -o ../../main.site/plugin/languages.mongodb.so