package jx3api

import "log/slog"

type Client struct {
	Opts *Options
}

type Options struct {
	Token  string
	Ticket string
}

func NewClient(opts *Options) *Client {
	if opts == nil {
		opts = &Options{}
	}

	if opts.Token == "" {
		slog.Info("The `token` parameter is not specified, only the free API can be used.")
	}

	return &Client{Opts: opts}
}
