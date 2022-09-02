.PHONY: buildimage dev shell build publish sync

buildimage:
	docker build -t "corso/website:latest" .

dev:
	docker run --rm -it -p 3000:3000 \
		-v ${PWD}:/usr/src/website corso/website gulp

shell:
	docker run --rm -it -p 3000:3000 \
		-v ${PWD}:/usr/src/website corso/website /bin/bash

build:
	docker run --rm -it \
		-v ${PWD}:/usr/src/website corso/website gulp build

publish: build
	docker run -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY \
		-e AWS_SESSION_TOKEN --rm -v ${PWD}:/usr/src/website corso/website \
		make sync

sync:
	aws s3 sync /usr/src/website/dist/ s3://corsobackup.io/ --exclude ".git/*" --delete
	aws cloudfront create-invalidation --distribution-id XXXXX --paths "/*"
