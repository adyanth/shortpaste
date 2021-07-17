FROM golang:alpine AS build

WORKDIR /usr/local/go/src/git.adyanth.site/adyanth/shortpaste/

RUN apk add --no-cache make build-base
COPY go.* ./
RUN go mod download

COPY *.go ./
COPY cmd cmd
COPY templates ./templates
RUN CGO_ENABLED=1 go build -o /out/ ./...

FROM alpine

WORKDIR /usr/local/bin/shortpaste/
COPY --from=build /out/shortpaste .

EXPOSE 8080

ENTRYPOINT [ "/shortpaste" ]
