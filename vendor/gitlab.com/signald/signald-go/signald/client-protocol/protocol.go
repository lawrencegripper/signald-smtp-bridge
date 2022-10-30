package client_protocol

import (
	"encoding/json"
)

type BasicResponse struct {
	ID        string          `json:"id,omitempty"`
	Type      string          `json:"type,omitempty"`
	ErrorType string          `json:"error_type,omitempty"`
	Error     json.RawMessage `json:"error,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
	Account   string          `json:"account,omitempty"`
}
