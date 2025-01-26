package sse

import "time"

var (
	RetryDuration int = int((time.Second * 3).Milliseconds())
)
