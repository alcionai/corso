.PHONY: buildimage dev shell build publish sync clean

WEBC := docker run --rm -it --init -p 5050:3000 -p 5051:3001 \
		-v ${PWD}:/usr/src/website corso/website:latest
ifeq (${CI},true)
# Cannot run in interactive/tty mode in CI
WEBC := docker run --rm --init -p 5050:3000 -p 5051:3001 \
		-v ${PWD}:/usr/src/website corso/website:latest
endif

buildimage: clean
	docker build -t "corso/website:latest" .

dev:
	$(WEBC) gulp

shell:
	$(WEBC) /bin/bash

build:
	$(WEBC) gulp build

publish: build
	docker run -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY \
		-e AWS_SESSION_TOKEN -e AWS_REGION \
		--rm -v ${PWD}:/usr/src/website corso/website:latest \
		make sync

sync:
	aws s3 sync /usr/src/website/dist/ s3://corsobackup.io/ --exclude ".git/*" --delete
	aws cloudfront create-invalidation --distribution-id E1W9NGI9YTVZ1A --paths "/*"

clean:
	rm -rf node_modules
