# syntax=docker/dockerfile:1

# Fetch
FROM golang:1.22.1 AS fetch-stage
COPY ./site /vts/site/
COPY go.mod go.sum ./cmd /vts/
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /vts
WORKDIR /vts
RUN ["templ", "generate"]

# Build
FROM golang:1.22.1 AS build-stage
COPY --from=generate-stage /vts /vts
WORKDIR /vts
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /vts/main /vts/cmd/

# Deploy
FROM alpine:latest AS deploy-stage
WORKDIR /vts
COPY --from=build-stage /vts/ ./

RUN apk add minify
RUN minify -o /vts/site/render/dist/tailwind.min.css /vts/site/render/dist/tailwind.css
RUN minify -o /vts/site/render/dist/events.min.js /vts/site/render/dist/events.js

ENV AWS_DEFAULT_REGION=eu-west-2
ENTRYPOINT ["/vts/main"]
EXPOSE 80
EXPOSE 443