# Golang base image
FROM golang:1.18 as go_builder
LABEL stage="hayasui_builder"
WORKDIR /hayasui
COPY . .
WORKDIR /hayasui/cmd/hayasui
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod vendor -o hayasui

# New stage from scratch
FROM alpine:3.16
LABEL stage="hayasui_binary"
COPY --from=go_builder /hayasui/cmd/hayasui/hayasui hayasui