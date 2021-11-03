FROM golang:1.17-alpine
WORKDIR /go/src

RUN apk update \
    apk add --no-cache git gcc musl-dev

RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/stretchr/testify/assert

COPY ./ ./
