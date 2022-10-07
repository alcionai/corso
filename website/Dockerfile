FROM ubuntu:22.04
LABEL MAINTAINER="Niraj Tolia"

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /usr/src/

RUN apt-get -y update && apt-get -y install curl zip make git emacs \
  && curl -fsSL https://deb.nodesource.com/setup_current.x | bash - \
  && apt-get -y install nodejs \
  && apt-get autoclean \
  && npm install -g gulp \
  && corepack enable \
  && node --version \
  && npm --version \
  && yarn --version \
  && gulp --version

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && /bin/bash /usr/src/aws/install

WORKDIR /usr/src
COPY package.json yarn.lock ./
RUN yarn install \
  && yarn cache clean \
  && rm -f package.json yarn.lock
ENV PATH /usr/src/node_modules/.bin:$PATH
ENV NODE_MODULES_PATH ../

WORKDIR /usr/src/website
