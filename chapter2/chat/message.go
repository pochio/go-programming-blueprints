package main

import (
	"time"
)

// message は１つのメッセージで表します
type message struct {
	Name    string
	Message string
	When    time.Time
}
