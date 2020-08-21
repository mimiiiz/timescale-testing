package model

import "time"

type Condition struct {
	Time        time.Time
	DeviceID    string
	Temperature uint
	humidity    uint
}
