FROM golang:1.20-alpine AS build

ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

WORKDIR /app
RUN mkdir -p /app/logs

RUN mkdir -p /app/bin
RUN mkdir -p /app/config
RUN mkdir -p /app/src/app
RUN mkdir -p /app/src/business
RUN mkdir -p /app/src/database
RUN mkdir -p /app/src/logger
RUN mkdir -p /app/src/model


COPY go.mod ./
COPY go.sum ./


COPY config/*.properties ./config/
COPY src/app/*.go ./src/app/
COPY src/business/*.go ./src/business/
COPY src/database/*.go ./src/database/
COPY src/logger/*.go ./src/logger/
COPY src/model/*.go ./src/model/
RUN go build -o bin ./...

#If not using sqlite3, choose scratch and set CGO_ENABLED to 0
FROM golang:1.20-alpine 
WORKDIR /app
COPY --from=build /app/bin/app /app/bin/app
COPY --from=build /app/config/* /app/config/
COPY --from=build /app/logs /app/logs
EXPOSE 8080
CMD ["/app/bin/app"]