FROM golang:1.17-alpine

RUN apk add build-base

RUN mkdir -p /etc/eth-service

WORKDIR /etc/eth-service

COPY ./cmd ./cmd

COPY ./internal ./internal

COPY ./go.mod ./

RUN go mod tidy

ARG WORKDIR=${WORKDIR}

WORKDIR ${WORKDIR}