package jx3api

import (
	"log/slog"
	"net/http"
	"os"
)

type Client struct {
	Opts *Options
}

type Options struct {
	Token      string
	Ticket     string
	Node       string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(opts *Options) *Client {
	if opts == nil {
		opts = &Options{}
	}

	if opts.Token == "" {
		token := os.Getenv("JX3API_TOKEN")
		if token != "" {
			opts.Token = token
		} else {
			slog.Info("The `token` parameter is not specified, only the free API can be used.")
		}
	}

	if opts.Ticket == "" {
		ticket := os.Getenv("JX3API_TICKET")
		if ticket != "" {
			opts.Ticket = ticket
		}
	}

	return &Client{Opts: opts}
}
