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
	"context"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	client_protocol "gitlab.com/signald/signald-go/signald/client-protocol"
)

const (
	defaultSocketPrefix = "/var/run"
	socketSuffix        = "/signald/signald.sock"

	ProductionServerUUID = "6e2eb5a8-5706-45d0-8377-127a816411a4"
	StagingServerUUID    = "97c17f0c-e53b-426f-8ffa-c052d4183f83"
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
	listeners  map[string]chan client_protocol.BasicResponse
	SocketPath string
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
	var d net.Dialer

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	addr := net.UnixAddr{Name: s.SocketPath, Net: "unix"}
	socket, err := d.DialContext(ctx, "unix", addr.String())
	if err != nil {
		return err
	}

	s.socket = socket
	return nil
}

func (s *Signald) Close() error {
	return s.socket.Close()
}

// Listen listens for events from signald
func (s *Signald) Listen(c chan client_protocol.BasicResponse) {
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

		if c != nil && !(msg.ID == "" && msg.Type == "version") {
			c <- msg
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

func (s *Signald) GetResponseListener(requestid string) chan client_protocol.BasicResponse {
	if s.listeners == nil {
		s.listeners = map[string]chan client_protocol.BasicResponse{}
	}
	c, ok := s.listeners[requestid]
	if !ok {
		c = make(chan client_protocol.BasicResponse)
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

func (s *Signald) readNext() (b client_protocol.BasicResponse, err error) {
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
