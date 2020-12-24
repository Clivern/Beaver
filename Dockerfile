FROM golang:1.15.2

ARG BEAVER_VERSION=2.0.0

ENV GO111MODULE=on

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs
RUN mkdir -p /app/var/storage
RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/Clivern/Beaver/releases/download/v${BEAVER_VERSION}/Beaver_${BEAVER_VERSION}_Linux_x86_64.tar.gz | tar xz
RUN rm LICENSE
RUN rm README.md
RUN mv Beaver beaver

COPY ./config.dist.yml /app/configs/

EXPOSE 8000

VOLUME /app/configs
VOLUME /app/var

RUN ./beaver version

CMD ["./beaver", "serve", "-c", "/app/configs/config.dist.yml"]