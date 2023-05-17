FROM node:alpine as node
COPY . /src
WORKDIR /src/nextjs/
RUN npm install
RUN cp next.config.build.js next.config.js
RUN npm run export

FROM golang:alpine as build-env
COPY . /src
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY --from=node /src/nextjs/out /src/nextjs/out
RUN GOOS=linux GOARCH=amd64 go build main.go

FROM alpine

LABEL maintainer="Mohammad Mahdi Afshar <me@mamad.dev>"
LABEL version="1.0"
LABEL description="A Simple Dockerized Golang + NextJS 13 App."
LABEL org.opencontainers.image.description="A Simple Dockerized Golang + NextJS 13 App."
LABEL org.opencontainers.image.authors="Mohammad Mahdi Afshar <me@mamad.dev>"
LABEL org.opencontainers.image.url="https://github.com/reloadlife/nextgo"

WORKDIR /
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=build-env /src/main /usr/local/bin/app
RUN chmod +x /usr/local/bin/app
ENTRYPOINT [ "/usr/local/bin/app" ]