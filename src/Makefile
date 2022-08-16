# This must match the version defined in .github/workflows/lint.yaml.
WANTED_LINT_VERSION := 1.46.2
LINT_VERSION := $(shell golangci-lint version | cut -d' ' -f4)
HAS_LINT := $(shell which golangci-lint)

INSTALL_LINT_PAGE := "https://golangci-lint.run/usage/install/"
BAD_LINT_MSG := "Missing golangci-lint version $(WANTED_LINT_VERSION). Visit $(INSTALL_LINT_PAGE) for instructions on how to install"

.PHONY: check-lint check-lint-version lint

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
