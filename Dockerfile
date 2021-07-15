FROM golang:alpine AS build

WORKDIR /usr/local/go/src/git.adyanth.site/adyanth/shortpaste/main

RUN apk add --no-cache make build-base
COPY main/go.mod .
COPY main/go.sum .
RUN go mod download

COPY main/*.go .

RUN CGO_ENABLED=1 go build -o /shortpaste

FROM alpine

COPY --from=build /shortpaste /

EXPOSE 8080

ENTRYPOINT [ "/shortpaste" ]
