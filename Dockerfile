# Step 1 Build
# FROM golang:1.12.0-alpine3.9 AS build
# RUN apk --no-cache add gcc g++ make
# RUN apk add git
# WORKDIR /go/src/app
# COPY . .
# ENV GO111MODULE=on
# ENV PORT=3000
# RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./*.go

#Step 2 Final
# FROM alpine:3.9
FROM golang:1.12.0-alpine3.9
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add gcc g++ make
RUN apk add git
# WORKDIR
# COPY --from=build /go/src/app/bin /go/bin
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV LANG C.UTF-8
ENV GO111MODULE=on
RUN apk --update add --no-cache bash git
RUN go get github.com/codegangsta/gin
EXPOSE 3000
WORKDIR /go/src/app
ENTRYPOINT gin --appPort 8080
