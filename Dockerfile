FROM golang:1.16.3-alpine3.13 as builder

RUN apk update && apk add gcc g++ musl-dev

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -ldflags="-w -s" -o movie-box .

FROM alpine:3.13.4

WORKDIR /app

COPY --from=builder /app/movie-box /usr/local/bin/

EXPOSE 3080

CMD [ "movie-box" , "server" ]

ENTRYPOINT [ "movie-box" , "get" ]
