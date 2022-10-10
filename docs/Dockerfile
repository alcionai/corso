FROM ubuntu:22.04
LABEL MAINTAINER="Niraj Tolia"

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -y update && apt-get -y install gpg emacs curl git make \
  && curl -fsSL https://deb.nodesource.com/setup_current.x | bash - \
  && apt-get -y install nodejs \
  && apt-get autoclean \
  && node --version \
  && npm --version \
  && cd /tmp && curl -O -L https://github.com/errata-ai/vale/releases/download/v2.20.1/vale_2.20.1_Linux_64-bit.tar.gz \ # NOTE: update in CI when updating
  && tar -xvzf vale_2.20.1_Linux_64-bit.tar.gz -C /usr/bin vale \
  && npm install -g markdownlint-cli@0.32.2 # NOTE: update in CI when updating

WORKDIR /usr/src
COPY package.json package-lock.json* ./
RUN npm ci \
  && npm cache clean --force \
  && rm -f package.json package-lock.json*
ENV PATH /usr/src/node_modules/.bin:$PATH

WORKDIR /usr/src/docs

CMD ["npm", "start", "--", "--host", "0.0.0.0"]
