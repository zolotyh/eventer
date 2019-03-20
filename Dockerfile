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
FROM golang:1.12.0-alpine3.9 AS build
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /usr/bin
# COPY --from=build /go/src/app/bin /go/bin
RUN go get -u github.com/gin-gonic/gin
EXPOSE 3000
ENTRYPOINT gin --appPort 8080
