package jx3api

import (
	"context"
	"io"
	"log/slog"

	"net/http"
)

func (c *Client) request(ctx context.Context, endpoint string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", BASE_URL+endpoint, body)
	if err != nil {
		slog.Error("request: requester build error: " + err.Error())
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	if c.Opts.Token != "" {
		slog.Debug("request: using token: " + c.Opts.Token)
		req.Header.Add("token", c.Opts.Token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("request: request error: " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("request: reading response body error: " + err.Error())
		return nil, err
	}

	return respBody, nil
}
