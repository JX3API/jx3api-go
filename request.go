package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

func (c *Client) request(ctx context.Context, endpoint string, params any) (json.RawMessage, error) {
	buf, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var payload map[string]json.RawMessage
	if err := json.Unmarshal(buf, &payload); err != nil {
		return nil, err
	}
	if payload == nil {
		payload = make(map[string]json.RawMessage)
	}
	for name, fallback := range map[string]string{"token": c.Opts.Token, "ticket": c.Opts.Ticket} {
		var value string
		if raw, ok := payload[name]; ok {
			if err := json.Unmarshal(raw, &value); err != nil {
				return nil, err
			}
		}
		if value == "" && fallback != "" {
			payload[name], _ = json.Marshal(fallback)
		}
	}
	buf, err = json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	baseURL := c.Opts.BaseURL
	if baseURL == "" {
		baseURL = BASE_URL
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+endpoint, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.Opts.Node != "" {
		req.Header.Set("node", c.Opts.Node)
	}

	client := c.Opts.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
	}
	return parseResponse(body)
}

func requestData[T any](ctx context.Context, c *Client, endpoint string, params any) (*T, error) {
	raw, err := c.request(ctx, endpoint, params)
	if err != nil {
		return nil, err
	}
	data := new(T)
	// The API can return an empty object for an empty list.
	_, custom := any(data).(json.Unmarshaler)
	if !custom && reflect.TypeFor[T]().Kind() == reflect.Slice && bytes.Equal(bytes.TrimSpace(raw), []byte("{}")) {
		return data, nil
	}
	if err := json.Unmarshal(raw, data); err != nil {
		return nil, err
	}
	return data, nil
}
