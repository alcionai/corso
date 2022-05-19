.PHONY: buildimage build dev shell check

buildimage:
	docker build -t "alcion/docs:latest" .

dev:
	docker run --rm -it \
		-p 3000:3000 \
		-v ${PWD}:/usr/src/docs alcion/docs

check:
	docker run --rm \
		-v ${PWD}:/usr/src/docs alcion/docs vale --glob='!**/development/*' docs
	docker run --rm \
		-v ${PWD}:/usr/src/docs alcion/docs markdownlint '**/*.md' --ignore styles/ --ignore src/ --ignore docs/development

shell:
	docker run --rm -it \
		-p 3000:3000 \
		-v ${PWD}:/usr/src/docs alcion/docs /bin/bash

build:
	docker run --rm \
		-v ${PWD}:/usr/src/docs alcion/docs npm run build
