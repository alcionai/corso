# This must match the version defined in .github/workflows/lint.yaml.
WANTED_LINT_VERSION := 1.45.2
LINT_VERSION := $(shell golangci-lint version | cut -d' ' -f4)
HAS_LINT := $(shell which golangci-lint)

INSTALL_LINT_PAGE := "https://golangci-lint.run/usage/install/"
BAD_LINT_MSG := "Missing golangci-lint version $(WANTED_LINT_VERSION). Visit $(INSTALL_LINT_PAGE) for instructions on how to install"

.PHONY: check-lint check-lint-version lint load-test

build:
	go build -o corso -ldflags \
	"-X 'github.com/alcionai/corso/src/cli/version.Version=$(shell git describe --exact-match --tags $(git rev-parse HEAD) 2>/dev/null || echo unreleased)-$(shell git rev-parse --short HEAD)'"

lint: check-lint-version
	golangci-lint run

check-lint-version: check-lint
	@if [ "$(LINT_VERSION)" != "$(WANTED_LINT_VERSION)" ]; then \
		echo >&2 $(BAD_LINT_MSG); \
		false; \
	fi

check-lint:
	@if [ -z "$(HAS_LINT)" ]; then \
		echo >&2 $(BAD_LINT_MSG); \
		false; \
	fi

build-otel-daemon:
	cd testfiles/otel_daemon; \
	docker build -t xray-daemon .

otel-daemon:
	results_dir=$$PWD/test_results; \
	cd ./testfiles/otel_daemon; \
	docker run \
		-d \
		-e AWS_REGION \
		-v ~/.aws/:/root/.aws/:ro \
		--name otel-daemon \
		-p 2000:2000/udp \
		--rm \
	xray-daemon \
		--local-mode \
		--log-level debug 

local-daemon:
	results_dir=$$PWD/test_results; \
	cd ./testfiles/otel_daemon; \
	docker run \
		--attach STDOUT \
		-e AWS_REGION \
		-v ~/.aws/:/root/.aws/:ro \
		--name otel-daemon \
		-p 2000:2000/udp \
		--rm \
	xray-daemon \
		--local-mode \
		--log-level debug 

# --net=host \

load-test:
	AWS_XRAY_NOOP_ID=False \
	CORSO_LOAD_TESTS=y \
	go test \
	-v \
	-count=1 \
	-timeout 1h \
	-blockprofile=block.prof \
	-cpuprofile=cpu.prof \
	-memprofile=mem.prof \
	-mutexprofile=mutex.prof \
	-trace=trace.out \
	-outputdir=test_results \
	./pkg/repository/repository_load_test.go