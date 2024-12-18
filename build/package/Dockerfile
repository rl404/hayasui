# Golang base image
FROM golang:1.23 AS go_builder
LABEL stage="hayasui_builder"
WORKDIR /hayasui
COPY . .
WORKDIR /hayasui/cmd/hayasui
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o hayasui

# New stage from scratch
FROM alpine:3.21
LABEL stage="hayasui_binary"
COPY --from=go_builder /hayasui/cmd/hayasui/hayasui hayasui