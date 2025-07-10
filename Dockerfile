
FROM golang:1.23.4-alpine AS builder

RUN apk --no-cache add gcc g++ make
RUN apk add git

WORKDIR /app
ADD . /app

RUN cd /app & go mod download
RUN cd /app & CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o topupservice main.go

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /app/topupservice /app
COPY --from=builder /app /app

EXPOSE 8080

ENTRYPOINT ./topupservice
