package logx

import (
	"time"
)

type Options struct {
	Level           string        `json:"level"`
	Caller          bool          `json:"caller"`
	Stacktrace      bool          `json:"stacktrace"`
	FlushInterval   time.Duration `json:"flush_interval"`
	DisableSampling bool          `json:"disable_sampling"`
	StdOut          bool          `json:"stdout"`
	File            string        `json:"file"`
	Size            int           `json:"size"`
	Age             int           `json:"age"`
	Backups         int           `json:"backups"`
	Compress        bool          `json:"compress"`
}

func DefaultOptions() *Options {
	return &Options{
		Level:         "INFO",
		Caller:        true,
		Stacktrace:    true,
		FlushInterval: 5 * time.Second,
		StdOut:        true,
		Age:           30,
		Size:          64 * 1024 * 1024,
		Backups:       64,
		Compress:      false,
	}
}
