FROM ubuntu:22.04
LABEL MAINTAINER="Niraj Tolia"

ARG DEBIAN_FRONTEND=noninteractive

# NOTE for lines 13,15: update in CI when updating
RUN apt-get -y update && apt-get -y install gpg emacs curl zip git make \
  && curl -fsSL https://deb.nodesource.com/setup_current.x | bash - \
  && apt-get -y install nodejs \
  && apt-get autoclean \
  && node --version \
  && npm --version \
  && arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/64-bit/) \
  && cd /tmp && curl -O -L https://github.com/errata-ai/vale/releases/download/v2.20.1/vale_2.20.1_Linux_${arch}.tar.gz \
  && tar -xvzf vale_2.20.1_Linux_${arch}.tar.gz -C /usr/bin vale \
  && npm install -g markdownlint-cli@0.32.2 \
  && curl "https://awscli.amazonaws.com/awscli-exe-linux-$(arch).zip" -o "awscliv2.zip" \
  && unzip awscliv2.zip && /bin/bash aws/install && rm -rf awscliv2.zip aws

WORKDIR /usr/src
COPY package.json package-lock.json* ./
RUN npm ci \
  && npm cache clean --force \
  && rm -f package.json package-lock.json*
ENV PATH /usr/src/node_modules/.bin:$PATH

WORKDIR /usr/src/website

CMD ["npm", "start", "--", "--host", "0.0.0.0"]
