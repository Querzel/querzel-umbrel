package main

import "fmt"

const (
	devs Command = "devs"
)

type JsonDevs struct {
	ASC                 int     `json:"ASC"`
	Name                string  `json:"Name"`
	Id                  int     `json:"ID"`
	Enabled             string  `json:"Enabled"`
	Status              string  `json:"Status"`
	Temperature         float32 `json:"Temperature"`
	MHSav               float32 `json:"MHS av"`
	MHS5s               float32 `json:"MHS 5s"`
	MHS1m               float32 `json:"MHS 1m"`
	MHS5m               float32 `json:"MHS 5m"`
	MHS15m              float32 `json:"MHS 15m"`
	Accepted            int     `json:"Accepted"`
	Rejected            int     `json:"Rejected"`
	HardwareErrors      int     `json:"Hardware Errors"`
	Utility             float32 `json:"Utility"`
	LastSharePool       int     `json:"Last Share Pool"`
	LastShareTime       int     `json:"Last Share Time"`
	TotalMH             float64 `json:"Total MH"`
	Diff1Work           int     `json:"Diff1 Work"`
	DifficultyAccepted  float64 `json:"Difficulty Accepted"`
	DifficultyRejected  float64 `json:"Difficulty Rejected"`
	LastShareDifficulty float64 `json:"Last Share Difficulty"`
	NoDevice            bool    `json:"No Device"`
	LastValidWork       int     `json:"Last Valid Work"`
	DeviceHardware      float32 `json:"Device Hardware%"`
	DeviceRejected      float32 `json:"Device Rejected%"`
	DeviceElapsed       int     `json:"Device Elapsed"`
}

type JsonDevsResult struct {
	JsonResult
	Items []JsonDevs `json:"DEVS,omitempty"`
}

func (res *JsonDevsResult) IsNotValid() bool {
	return res.HasInvalidStatus() ||
		res.Items == nil ||
		len(res.Items) <= 0
}

func cgminerDevs() (*JsonDevsResult, error) {
	command := devs
	if result, err := cgminerCommand(command); err != nil {
		return nil, fmt.Errorf("failed commnad=%s: %w", command, err)
	} else {
		if json, err := result.unmarshalDevs(); err != nil {
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

func (resp *RawResponse) unmarshalDevs() (*JsonDevsResult, error) {
	var jsonDevs JsonDevsResult
	if err := jsonUnmarshal(*resp, &jsonDevs); err != nil {
		return nil, err
	} else {
		return &jsonDevs, nil
	}
}
