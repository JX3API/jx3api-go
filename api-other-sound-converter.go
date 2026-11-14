package jx3api

import (
	"context"
	"encoding/json"
	"strconv"
)

type SoundConverterResponse = map[string]json.RawMessage

// SoundConverter 将文字转换为阿里云语音播放地址。
func (c *Client) SoundConverter(ctx context.Context, appkey string, access string, secret string, voice string, format string, sampleRate int, volume int, speechRate int, pitchRate int, text string) (*SoundConverterResponse, error) {
	params := &struct {
		Appkey     string `json:"appkey"`
		Access     string `json:"access"`
		Secret     string `json:"secret"`
		Voice      string `json:"voice,omitempty"`
		Format     string `json:"format,omitempty"`
		SampleRate int    `json:"sample_rate"`
		Volume     string `json:"volume,omitempty"`
		SpeechRate string `json:"speech_rate,omitempty"`
		PitchRate  string `json:"pitch_rate,omitempty"`
		Text       string `json:"text"`
	}{
		Appkey:     appkey,
		Access:     access,
		Secret:     secret,
		Voice:      voice,
		Format:     format,
		SampleRate: sampleRate,
		Volume:     strconv.Itoa(volume),
		SpeechRate: strconv.Itoa(speechRate),
		PitchRate:  strconv.Itoa(pitchRate),
		Text:       text,
	}
	return requestData[SoundConverterResponse](ctx, c, "/sound/converter", params)
}
