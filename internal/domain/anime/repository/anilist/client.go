package anilist

import (
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/verniy"
)

// Client contains functions for anilist.
type Client struct {
	client    *verniy.Client
	converter *md.Converter
}

// New to create new anilist.
func New() *Client {
	client := verniy.New()
	client.Http.Transport = newrelic.NewRoundTripper(http.DefaultTransport)
	return &Client{
		client:    client,
		converter: md.NewConverter("", true, nil),
	}
}
