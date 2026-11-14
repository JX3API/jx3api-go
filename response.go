package jx3api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status int             `json:"status"`
	Detail json.RawMessage `json:"detail"`
}

type APIError struct {
	Status int
	Detail string
}

func (e *APIError) Error() string {
	return e.Detail
}

func parseResponse(body []byte) (json.RawMessage, error) {
	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	if resp.Status == 0 || len(resp.Detail) == 0 {
		return nil, fmt.Errorf("invalid API response: status and detail are required")
	}
	if resp.Status != http.StatusOK {
		var message string
		if err := json.Unmarshal(resp.Detail, &message); err != nil {
			return nil, fmt.Errorf("invalid API error detail: %w", err)
		}
		return nil, &APIError{Status: resp.Status, Detail: message}
	}
	return resp.Detail, nil
}

func unmarshalRecords[T any](raw []byte, records *[]T) error {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 {
		return fmt.Errorf("empty API detail")
	}
	if raw[0] == '{' {
		if bytes.Equal(raw, []byte("{}")) {
			*records = nil
			return nil
		}
		var record T
		if err := json.Unmarshal(raw, &record); err != nil {
			return err
		}
		*records = []T{record}
		return nil
	}
	return json.Unmarshal(raw, records)
}
