# Golang base image
FROM golang:1.27.3 AS go_builder
LABEL stage="hayasui_builder"
WORKDIR /hayasui
COPY . .
WORKDIR /hayasui/cmd/hayasui
RUN CGO_ENABLED=0 go build -mod=vendor -ldflags="-s -w" -o hayasui

# New stage from scratch
FROM alpine:3.24.4
LABEL stage="hayasui_binary"
RUN adduser -D -u 10001 hayasui
COPY --from=go_builder /hayasui/cmd/hayasui/hayasui /usr/local/bin/hayasui
WORKDIR /home/hayasui
USER hayasui
ENTRYPOINT ["/usr/local/bin/hayasui"]
