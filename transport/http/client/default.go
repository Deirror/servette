package client

import (
	"time"
)

var defaultCfg = Config{
    ReadTimeout:  15 * time.Second,  // total request time
    WriteTimeout: 5 * time.Second,   // connection/dial timeout
    IdleTimeout:  90 * time.Second,  // keep-alive idle
}

func DefaultConfig() Config {
    return defaultCfg
}
