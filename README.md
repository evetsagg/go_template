# Go Microservice Template (WIP)

A Go Microservice template with sample logging (using zap), daos, docker files, etc.

## General build and run instructions

To build

<code>go build -o bin ./...</code>

To run app using the binary built

<code>./bin/app</code>

## Using Docker

To build a docker image

<code>docker build -t <image_name>:<image_tag> .</code>

To run the built docker image

<code>docker run -d --name <container_name> -p 8080:8080 <image_name>:<image_tag</code>

To use Docker-compose

<code>docker-compose up -d</code>

## TODO

- Sample Rest API client
- Prometheus integration
- Grafana integration
- Use Gin framework
