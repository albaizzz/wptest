package models

import "encoding/xml"

type DeviceModel struct {
	Device string  `xml:"device"`
	Value  float64 `xml:"value"`
}

type DeviceArray1 struct {
	Devices []Devices `xml:",any"`
}
type Devices struct {
	XMLName xml.Name `xml:""`
	Value   string   `xml:",chardata"`
}

type DeviceRequest struct {
	ID          int           `xml:"id"`
	RecordTime  int64         `xml:"record_time"`
	Devicemodel []DeviceModel `xml:"data>element"`
	Devices     DeviceArray   `xml:"devices"`
}
