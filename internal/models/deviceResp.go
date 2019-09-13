package models

type DeviceResp struct {
	Value    float64 `json:"value"`
	DateTime string  `json:"datetime"`
	Unit     string  `json:"unit"`
}
