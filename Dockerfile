# syntax=docker/dockerfile:1.2
FROM golang:1.16-buster as builder

WORKDIR /build
COPY . .
RUN --mount=type=cache,target=/go/pkg --mount=type=cache,target=/root/.cache/go-build make build

FROM debian:buster
LABEL org.opencontainers.image.source https://github.com/lawrencegripper/signald-smtp-bridge  

WORKDIR /app
COPY --from=builder /build/signald-smtp-bridge /app

ENTRYPOINT [ "/app/signald-smtp-bridge" ]
