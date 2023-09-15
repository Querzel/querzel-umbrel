package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Command string

type JsonStatus struct {
	Status      string `json:"STATUS"`
	When        int    `json:"When"`
	Code        int    `json:"Code"`
	Message     string `json:"Msg"`
	Description string `json:"Description"`
}

func (c *Command) ToJsonString() string {
	return fmt.Sprintf(`{"command": "%s"}`, *c)
}

type JsonResultValidator interface {
	IsValid() bool
	IsNotValid() bool
}

type JsonStatusValidator interface {
	HasInvalidStatus() bool
}

type JsonResult struct {
	Statuses []JsonStatus `json:"STATUS"`
	Id       int          `json:"id"`
	JsonResultValidator
	JsonStatusValidator
}

func (res *JsonResult) IsValid() bool {
	return !res.IsNotValid()
}

func (res *JsonResult) HasInvalidStatus() bool {
	return res.Statuses == nil ||
		len(res.Statuses) <= 0 ||
		res.Statuses[0].Status != "S"
}

func jsonUnmarshal(data []byte, v interface{}) error {
	rawBytes := bytes.Trim(data, " \t\r\n\000")
	return json.Unmarshal(rawBytes, v)
}
