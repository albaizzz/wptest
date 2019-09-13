package models

import (
	"time"
)

type DeviceArray struct {
	Devices []Devices `json:",any"`
}

type Device struct {
	ID        int       `json:"id"`
	Device    string    `json:"device"`
	Value     float64   `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}
