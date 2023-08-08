.PHONY: buildimage build dev shell check genclidocs _validatemdgen publish sync

GO_VERSION := 1.19
CORSO_BUILD_DIR := /tmp/.corsobuild
CORSO_BUILD_CACHE := ${CORSO_BUILD_DIR}/cache
CORSO_BUILD_MOD := ${CORSO_BUILD_DIR}/mod
CORSO_BUILD_BIN := ${CORSO_BUILD_DIR}/bin
CORSO_REPO := /go/src/github.com/alcionai/corso
CORSO_LOCAL_PATH := $(shell git rev-parse --show-toplevel)
GIT_SHA := $(shell git rev-parse --short HEAD)
WEBC := docker run --rm -it -p 3000:3000 -v ${PWD}:/usr/src/website --env CORSO_VERSION=unreleased-${GIT_SHA} corso/website
CBASE :=  docker run --rm -it \
		-v ${CORSO_LOCAL_PATH}:${CORSO_REPO} -v ${CORSO_BUILD_DIR}:${CORSO_BUILD_DIR} \
		--env GOCACHE=${CORSO_BUILD_CACHE} --env GOMODCACHE=${CORSO_BUILD_MOD} --env GOTMPDIR=${CORSO_BUILD_DIR} \
		--workdir ${CORSO_REPO}/src
GOC :=  ${CBASE} golang:${GO_VERSION}
GOBASHC :=  ${CBASE} --entrypoint bash golang:${GO_VERSION}
MDGEN_SRC := ${CORSO_REPO}/src/cmd/mdgen/mdgen.go
MDGEN_BINARY := ${CORSO_BUILD_BIN}/mdgen
CLI_DOCS := ${CORSO_REPO}/website/docs/cli

buildimage:
	docker build -t "corso/website:latest" .

dev: genclidocs
	$(WEBC) npm start -- --host 0.0.0.0

noclidev:
	$(WEBC) npm start -- --host 0.0.0.0

VALE_TARGET ?= docs README.md blog

check: genclidocs
	$(WEBC) vale $(VALE_TARGET)
	$(WEBC) markdownlint '**/*.md' --ignore styles/ --ignore src/ --ignore node_modules/

localcheck: genclidocs
	vale $(VALE_TARGET)
	markdownlint '**/*.md' --ignore styles/ --ignore src/ --ignore node_modules/

dockershell:
	$(WEBC) bash

build: genclidocs
	$(WEBC) npm run build

genclidocs: _validatemdgen ${MDGEN_BINARY}
	@echo 'Auto-generating Corso CLI docs...'
	$(WEBC) rm -rf docs/cli
	$(GOC) ${MDGEN_BINARY} --cli-folder ${CLI_DOCS}

_validatemdgen: # in case we have a different architecture
	@echo 'Verifying dependencies...'
	$(GOBASHC) -c "${MDGEN_BINARY} --help >/dev/null || rm -rf ${MDGEN_BINARY}"

${MDGEN_BINARY}: $(shell find ${CORSO_LOCAL_PATH}/src -type f -name *.go) $(shell find ${CORSO_LOCAL_PATH}/src -type d )
	@echo 'Re-building Corso CLI docs auto-gen tooling...'
		$(GOC) go mod download 
		$(GOC) go build -o ${MDGEN_BINARY} ${MDGEN_SRC}

clean:
	$(WEBC) rm -rf docs/cli build node_modules
	$(GOC) rm -rf ${CORSO_BUILD_DIR}/*

publish: clean build
	docker run -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY \
		-e AWS_SESSION_TOKEN -e AWS_REGION \
		--rm -v ${PWD}:/usr/src/website corso/website:latest \
	make sync

sync:
	aws s3 sync /usr/src/website/build/ s3://corsobackup.io/ --delete
	aws cloudfront create-invalidation --distribution-id E1W9NGI9YTVZ1A --paths "/*"
