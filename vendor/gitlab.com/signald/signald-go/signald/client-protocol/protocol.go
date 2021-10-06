package client_protocol

import (
	"encoding/json"
)

type BasicResponse struct {
	ID        string
	Type      string
	ErrorType string
	Error     json.RawMessage
	Data      json.RawMessage
}
