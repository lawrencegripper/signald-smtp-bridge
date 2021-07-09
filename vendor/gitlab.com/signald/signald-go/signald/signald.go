// Copyright © 2021 Finn Herzfeld <finn@janky.solutions>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package signald

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"gitlab.com/signald/signald-go/signald/client-protocol/v0"
)

const (
	defaultSocketPrefix = "/var/run"
	socketSuffix        = "/signald/signald.sock"
)

var (
	debugSignaldIO, _ = strconv.ParseBool(os.Getenv("DEBUG_SIGNALD_IO"))
	xdgRuntimeDir     = os.Getenv("XDG_RUNTIME_DIR")
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Signald is a connection to a signald instance.
type Signald struct {
	socket     net.Conn
	listeners  map[string]chan BasicResponse
	SocketPath string
}

type BasicResponse struct {
	ID    string
	Type  string
	Error json.RawMessage
	Data  json.RawMessage
}

type UnexpectedError struct {
	Message string
}

// Connect connects to the signad socket
func (s *Signald) Connect() error {
	if s.SocketPath != "" {
		return s.connect()
	}

	s.SocketPath = xdgRuntimeDir + socketSuffix
	err := s.connect()
	if err != nil {
		_, ok := err.(net.Error)
		if ok {
			s.SocketPath = defaultSocketPrefix + socketSuffix
			err = s.connect()
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (s *Signald) connect() error {
	socket, err := net.Dial("unix", s.SocketPath)
	if err != nil {
		return err
	}
	s.socket = socket
	return nil
}

// Listen listens for events from signald
func (s *Signald) Listen(c chan v0.LegacyResponse) {
	for {
		msg, err := s.readNext()
		if err == io.EOF {
			log.Println("signald-go: socket disconnected!")
			if c != nil {
				close(c)
			}
			return
		}

		if msg.Type == "unexpected_error" {
			var errorResponse UnexpectedError
			if err := json.Unmarshal(msg.Data, &errorResponse); err != nil {
				log.Println("signald-go: Error unmarshaling error response:", err.Error())
				continue
			}
			log.Println("signald-go: Unexpected error", errorResponse.Message)
			continue
		}

		if subscribers, ok := s.listeners[msg.ID]; ok {
			subscribers <- msg
		}

		if c != nil {
			legacyResponse := v0.LegacyResponse{ID: msg.ID, Type: msg.Type}
			if err := json.Unmarshal(msg.Data, &legacyResponse.Data); err != nil {
				log.Println("signald-go receive error: ", err)
			} else {
				c <- legacyResponse
			}
		}
	}
}

func (s *Signald) RawRequest(request interface{}) error {
	if debugSignaldIO {
		buffer := bytes.Buffer{}
		if err := json.NewEncoder(&buffer).Encode(request); err == nil {
			log.Println("[to signald]", strings.TrimSpace(buffer.String()))
		}
	}
	return json.NewEncoder(s.socket).Encode(request)
}

func (s *Signald) GetResponseListener(requestid string) chan BasicResponse {
	if s.listeners == nil {
		s.listeners = map[string]chan BasicResponse{}
	}
	c, ok := s.listeners[requestid]
	if !ok {
		c = make(chan BasicResponse)
		s.listeners[requestid] = c
	}
	return c
}

func (s *Signald) CloseResponseListener(requestid string) {
	listener, ok := s.listeners[requestid]
	if !ok {
		return
	}
	close(listener)
	delete(s.listeners, requestid)
}

func (s *Signald) readNext() (b BasicResponse, err error) {
	if debugSignaldIO {
		buffer := bytes.Buffer{}
		err = json.NewDecoder(io.TeeReader(s.socket, &buffer)).Decode(&b)
		log.Println("[from signald]", strings.TrimSpace(buffer.String()))
	} else {
		err = json.NewDecoder(s.socket).Decode(&b)
	}
	if err != nil {
		log.Println("signald-go: error decoding message from signald:", err)
		return
	}
	return
}

func (b BasicResponse) GetError() error {
	if b.Error == nil {
		return nil
	}
	return fmt.Errorf("signald error: %s", string(b.Error))
}
