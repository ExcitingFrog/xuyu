.PHONY: help
help:
	@echo 'Management commands for this service:'
	@echo
	@echo 'Usage:'
	@echo 'make bin                 Install all dependency bins'
	@echo 'make buf             	Init GRpc proto(install buf fisrt [go buf](https://github.com/bufbuild/buf))'
	@echo 'make mocks             	Generate test code(install buf fisrt [go mockery](https://vektra.github.io/mockery/latest/installation/#github-release))'
	@echo

.PHONY: buf
buf: buf-clean buf-generate

.PHONY: buf-clean
buf-clean:
	rm -rf proto/gen
	mkdir -p proto/gen

.PHONY: buf-generate
buf-generate:
	buf build -o lock.json
	buf generate

.PHONY: mocks
mocks:
	mockery --dir internal --all --output internal/mocks
