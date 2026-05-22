package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SoundConverterResponse struct {
	Text     string `json:"text"`
	Token    string `json:"token"`
	URL      string `json:"url"`
}

func (c *Client) SoundConverter(ctx context.Context, appkey string, access string, secret string, voice string, format string, sampleRate int, volume int, speechRate int, pitchRate int, text string) (*SoundConverterResponse, error) {
	params := &struct {
		Appkey    string `json:"appkey"`
		Access    string `json:"access"`
		Secret    string `json:"secret"`
		Voice     string `json:"voice"`
		Format    string `json:"format"`
		SampleRate int    `json:"sample_rate"`
		Volume    int    `json:"volume"`
		SpeechRate int    `json:"speech_rate"`
		PitchRate  int    `json:"pitch_rate"`
		Text      string `json:"text"`
	}{
		Appkey:    appkey,
		Access:    access,
		Secret:    secret,
		Voice:     voice,
		Format:    format,
		SampleRate: sampleRate,
		Volume:    volume,
		SpeechRate: speechRate,
		PitchRate:  pitchRate,
		Text:      text,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SoundConverter: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/sound/converter", body)
	if err != nil {
		slog.Error("SoundConverter: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SoundConverter: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(SoundConverterResponse)

	if resp.Msg != "success" {
		slog.Error("SoundConverter: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SoundConverter: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
