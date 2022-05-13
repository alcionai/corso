FROM ubuntu:22.04
LABEL MAINTAINER="Niraj Tolia"

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /usr/src/docs

RUN apt-get -y update && apt-get -y install gpg emacs curl \
  && curl -fsSL https://deb.nodesource.com/setup_current.x | bash - \
  && apt-get -y install nodejs \
  && apt-get autoclean \
  && node --version \
  && npm --version

CMD ["npm", "start", "--", "--host", "0.0.0.0"]
