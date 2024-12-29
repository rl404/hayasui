package anilist

import (
	"net/http"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/verniy"
)

// Client contains functions for anilist.
type Client struct {
	client    *verniy.Client
	converter *converter.Converter
}

// New to create new anilist.
func New() *Client {
	client := verniy.New()
	client.Http.Transport = newrelic.NewRoundTripper(http.DefaultTransport)
	return &Client{
		client: client,
		converter: converter.NewConverter(converter.WithPlugins(
			base.NewBasePlugin(),
			commonmark.NewCommonmarkPlugin(),
		)),
	}
}
