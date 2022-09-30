.PHONY: buildimage build dev shell check genclidocs

# 
CORSO_BUILD_DIR := /tmp/.corsobuild
CORSO_BUILD_CACHE := ${CORSO_BUILD_DIR}/cache
CORSO_BUILD_MOD := ${CORSO_BUILD_DIR}/mod
CORSO_BUILD_BIN := ${CORSO_BUILD_DIR}/bin

CORSO_REPO := ${PWD}/..
CORSO_REPO_CONTAINER := /go/src/github.com/alcionai/corso

MDGEN_SRC_CONTAINER := ${CORSO_REPO_CONTAINER}/src/cmd/mdgen/mdgen.go
MDGEN_BINARY := ${CORSO_BUILD_BIN}/mdgen
CLI_DOCS_CONTAINER := ${CORSO_REPO_CONTAINER}/docs/docs/cli


buildimage:
	docker build -t "alcion/docs:latest" .

dev: genclidocs
	docker run --rm -it \
		-p 3000:3000 \
		-v ${PWD}:/usr/src/docs alcion/docs

VALE_TARGET ?= docs README.md

dockercheck: genclidocs
	docker run --rm \
		-v ${PWD}:/usr/src/docs alcion/docs vale $(VALE_TARGET)
	docker run --rm \
	 	-v ${PWD}:/usr/src/docs alcion/docs markdownlint '**/*.md' --ignore styles/ --ignore src/ --ignore node_modules/

check: genclidocs
	vale $(VALE_TARGET)
	markdownlint '**/*.md' --ignore styles/ --ignore src/ --ignore node_modules/

shell:
	docker run --rm -it \
		-p 3000:3000 \
		-v ${PWD}:/usr/src/docs alcion/docs /bin/bash

build: genclidocs
	docker run --rm \
		-v ${PWD}:/usr/src/docs alcion/docs npm run build

genclidocs: ${MDGEN_BINARY}
	@echo 'Auto-generating Corso CLI docs...'
	rm -rf docs/cli
	docker run --rm -it \
		-v ${CORSO_REPO}:${CORSO_REPO_CONTAINER} -v ${CORSO_BUILD_DIR}:${CORSO_BUILD_DIR} \
		--workdir ${CORSO_REPO_CONTAINER}/src \
		--entrypoint ${MDGEN_BINARY} \
		golang:1.18 \
		--cli-folder ${CLI_DOCS_CONTAINER}

${MDGEN_BINARY}: $(shell find ../src -type f -name *.go) $(shell find ../src -type d )
	@echo 'Re-building Corso CLI docs auto-gen tooling...'
	docker run --rm -it \
		-v ${CORSO_REPO}:${CORSO_REPO_CONTAINER} -v ${CORSO_BUILD_DIR}:${CORSO_BUILD_DIR} \
		--env GOCACHE=${CORSO_BUILD_CACHE} --env GOMODCACHE=${CORSO_BUILD_MOD} --env GOTMPDIR=${CORSO_BUILD_DIR} \
		--workdir ${CORSO_REPO_CONTAINER}/src \
		--entrypoint /usr/local/go/bin/go \
		golang:1.18 \
		mod download 
	docker run --rm -it \
		-v ${CORSO_REPO}:${CORSO_REPO_CONTAINER} -v ${CORSO_BUILD_DIR}:${CORSO_BUILD_DIR} \
		--env GOCACHE=${CORSO_BUILD_CACHE} --env GOMODCACHE=${CORSO_BUILD_MOD} --env GOTMPDIR=${CORSO_BUILD_DIR} \
		--workdir ${CORSO_REPO_CONTAINER}/src \
		--entrypoint /usr/local/go/bin/go \
		golang:1.18 \
		build -o ${MDGEN_BINARY} ${MDGEN_SRC_CONTAINER}
