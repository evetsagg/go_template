FROM golang:1.20-alpine AS build
WORKDIR /app
RUN mkdir -p /app/logs
RUN mkdir -p /app/bin
RUN mkdir -p /app/src/app
RUN mkdir -p /app/src/business
RUN mkdir -p /app/src/database
RUN mkdir -p /app/src/logger
RUN mkdir -p /app/src/model

COPY go.mod ./
COPY go.sum ./
RUN go mod download
#Can just build locally and then copy the binary only if same architecture (x86, arm, etc.)
COPY src/app/*.go ./src/app/
COPY src/business/*.go ./src/business/
COPY src/database/*.go ./src/database/
COPY src/logger/*.go ./src/logger/
COPY src/model/*.go ./src/model/
#COPY  src/app/*.go ./
#COPY src/package2/*.go ./


#RUN go build -o go_rest
RUN go build -o bin ./...
#FROM scratch
#COPY --from=build /compose/hello-docker/backend /usr/local/bin/backend
EXPOSE 8080
CMD ["/app/bin/app"]