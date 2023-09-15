package main

import (
	"fmt"
)

const (
	lcd Command = "lcd"
)

type JsonLcd struct {
	Elapsed             int     `json:"Elapsed"`
	GHSAvg              float32 `json:"GHS av"`
	GHS5m               float32 `json:"GHS 5m"`
	GHS5s               float32 `json:"GHS 5s"`
	Temperature         float32 `json:"Temperature"`
	LastShareDifficulty float64 `json:"Last Share Difficulty"`
	LastShareTime       int     `json:"Last Share Time"`
	BestShare           int     `json:"BestShare"`
	LastValidWork       int     `json:"Last Valid Work"`
	FoundBlocks         int     `json:"Found Blocks"`
	CurrentPool         string  `json:"Current Pool"`
	User                string  `json:"User"`
}

type JsonLcdResult struct {
	*JsonResult
	Items []JsonLcd `json:"LCD,omitempty"`
}

func (res *JsonLcdResult) IsNotValid() bool {
	return res.HasInvalidStatus() ||
		res.Items == nil ||
		len(res.Items) <= 0
}

func cgminerLcd() (*JsonLcdResult, error) {
	command := lcd
	if result, err := cgminerCommand(command); err != nil {
		return nil, fmt.Errorf("failed commnad=%s: %w", command, err)
	} else {
		if json, err := result.unmarshalLcd(); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON, commnad=%s: %w payload=%s", command, err, string(*result))
		} else {
			if json.IsNotValid() {
				return nil, fmt.Errorf("failed to detect cgminer status, commnad=%s: %v, json=%+v\n", command, err, *json)
			} else {
				return json, nil
			}
		}
	}
}

func (resp *RawResponse) unmarshalLcd() (*JsonLcdResult, error) {
	var jsonLcd JsonLcdResult
	if err := jsonUnmarshal(*resp, &jsonLcd); err != nil {
		return nil, err
	} else {
		return &jsonLcd, nil
	}
}
