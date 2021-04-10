# Golang base image
FROM golang:1.15 as go_builder
LABEL stage="hayasui_builder"
WORKDIR /go/src/github.com/rl404/hayasui
COPY . .
WORKDIR /go/src/github.com/rl404/hayasui/cmd/hayasui
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod vendor -o hayasui

# New stage from scratch
FROM alpine:3.13
COPY --from=go_builder /go/src/github.com/rl404/hayasui/cmd/hayasui/hayasui hayasui