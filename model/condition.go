package model

import "time"

type Condition struct {
	Time        time.Time
	DeviceID    string
	Temperature float64
	Humidity    float64
}
