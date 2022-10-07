.PHONY: buildimage dev shell build publish sync

WEBC := docker run --rm -it --init -p 5050:3000 -p 5051:3001 \
		-v ${PWD}:/usr/src/website corso/website
ifeq (${CORSO_USE_DOCKER},-1)
	WEBC :=
endif

buildimage:
	docker build -t "corso/website:latest" .

dev:
	$(WEBC) gulp

shell:
	$(WEBC) /bin/bash

build:
	$(WEBC) gulp build

publish: build
	docker run -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY \
		-e AWS_SESSION_TOKEN --rm -v ${PWD}:/usr/src/website corso/website \
		make sync

sync:
	aws s3 sync /usr/src/website/dist/ s3://corsobackup.io/ --exclude ".git/*" --delete
	aws cloudfront create-invalidation --distribution-id XXXXX --paths "/*"
