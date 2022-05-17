FROM ubuntu:22.04
LABEL MAINTAINER="Niraj Tolia"

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -y update && apt-get -y install gpg emacs curl \
  && curl -fsSL https://deb.nodesource.com/setup_current.x | bash - \
  && apt-get -y install nodejs \
  && apt-get autoclean \
  && node --version \
  && npm --version

WORKDIR /usr/src
COPY package.json package-lock.json* ./
RUN npm ci && npm cache clean --force && rm -f package.json package-lock.json*
ENV PATH /usr/src/node_modules/.bin:$PATH

WORKDIR /usr/src/docs

CMD ["npm", "start", "--", "--host", "0.0.0.0"]
