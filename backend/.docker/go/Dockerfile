FROM golang:1.18-alpine

RUN apk update
RUN apk add --update --no-cache make alpine-sdk mysql-client bash git openssh wget

RUN mkdir /app

WORKDIR /app
