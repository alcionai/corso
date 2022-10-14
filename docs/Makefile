.PHONY: buildimage build dev shell check genclidocs _validatemdgen

# Specify `CORSO_USE_DOCKER=-1` to skip using Docker for builds
CORSO_BUILD_DIR := /tmp/.corsobuild
CORSO_BUILD_CACHE := ${CORSO_BUILD_DIR}/cache
CORSO_BUILD_MOD := ${CORSO_BUILD_DIR}/mod
CORSO_BUILD_BIN := ${CORSO_BUILD_DIR}/bin
CORSO_REPO := /go/src/github.com/alcionai/corso
CORSO_LOCAL_PATH := $(shell git rev-parse --show-toplevel)
DOCSC := docker run --rm -it -p 3000:3000 -v ${PWD}:/usr/src/docs alcion/docs
DOCSE := ${DOCSC} # for enforcing docker container as environment
GOC :=  docker run --rm -it \
		-v ${CORSO_LOCAL_PATH}:${CORSO_REPO} -v ${CORSO_BUILD_DIR}:${CORSO_BUILD_DIR} \
		--env GOCACHE=${CORSO_BUILD_CACHE} --env GOMODCACHE=${CORSO_BUILD_MOD} --env GOTMPDIR=${CORSO_BUILD_DIR} \
		--workdir ${CORSO_REPO}/src \
		golang:1.18
ifeq (${CORSO_USE_DOCKER},-1)
	DOCSC :=
	GOC := cd ${CORSO_LOCAL_PATH}/src &&
    CORSO_REPO := ${CORSO_LOCAL_PATH}
endif

MDGEN_SRC := ${CORSO_REPO}/src/cmd/mdgen/mdgen.go
MDGEN_BINARY := ${CORSO_BUILD_BIN}/mdgen
CLI_DOCS := ${CORSO_REPO}/docs/docs/cli

buildimage:
	docker build -t "alcion/docs:latest" .

dev: genclidocs
	$(DOCSC) npm start -- --host 0.0.0.0

VALE_TARGET ?= docs README.md

check: genclidocs
	$(DOCSC) vale $(VALE_TARGET)
	$(DOCSC) markdownlint '**/*.md' --ignore styles/ --ignore src/ --ignore node_modules/

dockershell:
	$(DOCSE) /bin/bash

build: genclidocs
	$(DOCSC) npm run build

genclidocs: ${MDGEN_BINARY}
	@echo 'Auto-generating Corso CLI docs...'
	$(DOCSC) rm -rf docs/cli
	$(GOC) ${MDGEN_BINARY} --cli-folder ${CLI_DOCS}

_validatemdgen: # in case we have a different architecture
	${MDGEN_BINARY} --help >/dev/null || rm -rf ${MDGEN_BINARY}

${MDGEN_BINARY}: $(shell find ${CORSO_LOCAL_PATH}/src -type f -name *.go) $(shell find ${CORSO_LOCAL_PATH}/src -type d ) _validatemdgen
	@echo 'Re-building Corso CLI docs auto-gen tooling...'
		$(GOC) go mod download 
		$(GOC) go build -o ${MDGEN_BINARY} ${MDGEN_SRC}

clean:
	$(DOCSC) rm -rf docs/cli
	$(DOCSC) rm -rf ${CORSO_BUILD_DIR}