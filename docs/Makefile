.PHONY: buildimage

buildimage:
	docker build -t "alcion/docs:latest" .

dev:
	docker run --rm -it \
		-p 3000:3000 \
		-v ${PWD}:/usr/src/docs alcion/docs

shell:
	docker run --rm -it \
		-p 3000:3000 \
		-v ${PWD}:/usr/src/docs alcion/docs /bin/bash

build:
	docker run --rm \
		-v ${PWD}:/usr/src/docs alcion/docs npm run build

