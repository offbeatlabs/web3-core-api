package dto

import "time"

// RestError proxy struct to be used with the swagger library
type RestError struct {
	ErrError   string      `json:"error,omitempty"`
	ErrMessage interface{} `json:"message,omitempty"`
	Timestamp  time.Time   `json:"timestamp,omitempty"`
}
