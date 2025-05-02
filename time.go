package main

import (
	"time"
)

func current_time() float64 {
	return float64(time.Now().UnixNano()) / 1e9
}
