.PHONY: buildimage

buildimage:
	docker build -t "corso/website:latest" .

dev:
	docker run -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY \
		-e AWS_SESSION_TOKEN --rm -it \
		-v ${PWD}:/usr/src/website corso/website /bin/bash

publish:
	docker run -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY \
		-e AWS_SESSION_TOKEN --rm -v ${PWD}:/usr/src/website corso/website \
		make sync

sync:
	aws s3 sync /usr/src/website/dist/ s3://corsobackup.io/ --exclude ".git/*" --delete
	aws cloudfront create-invalidation --distribution-id ESFSMEIU1TDEN --paths "/*"
