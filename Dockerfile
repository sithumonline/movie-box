FROM golang:1.16.3-alpine3.13 as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o movie-box .

FROM alpine:3.13.4

WORKDIR /app

COPY --from=builder /app/movie-box /usr/local/bin/

ENTRYPOINT [ "movie-box" , "get" ]
