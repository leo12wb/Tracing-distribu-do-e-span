FROM alpine:latest AS timezone_build
RUN apk --no-cache add tzdata ca-certificates  

FROM golang:1.21.6-alpine3.18 AS builder

RUN apk --no-cache add tzdata ca-certificates

ADD . /go/server

WORKDIR /go/server

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN /go/bin/swag init -g pkg/adapter/http/rest/rest.go -o pkg/adapter/http/rest/docs

RUN mkdir deploy
RUN go clean --modcache
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o go_app cmd/main.go 
RUN mv go_app ./deploy/go_app
RUN mv config.json ./deploy/config.json
RUN mv pkg/adapter/http/rest/docs/ ./deploy/docs

FROM scratch AS production

COPY --from=timezone_build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=timezone_build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/server/deploy /server/

WORKDIR /server

ENTRYPOINT  ["./go_app", "serve"]