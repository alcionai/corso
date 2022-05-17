FROM ubuntu:22.04
LABEL MAINTAINER="Niraj Tolia"

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /usr/src/docs

RUN apt-get -y update && apt-get -y install gpg emacs curl \
  && curl -fsSL https://deb.nodesource.com/setup_current.x | bash - \
  && apt-get -y install nodejs \
  && apt-get autoclean \
  && node --version \
  && npm --version \
  && cd /tmp && curl -O -L https://github.com/errata-ai/vale/releases/download/v2.15.4/vale_2.15.4_Linux_64-bit.tar.gz \
  && tar -xvzf vale_2.15.4_Linux_64-bit.tar.gz -C /usr/bin vale

CMD ["npm", "start", "--", "--host", "0.0.0.0"]
