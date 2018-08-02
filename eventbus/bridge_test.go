// Copyright 2016 Julien Ponge
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventbus

import (
	"encoding/json"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageRegister(t *testing.T) {

	listener, err := net.Listen("tcp", "localhost:7000")
	if err != nil {
		t.Error("Starting server failed", err)
	}
	defer listener.Close()

	sendMsg := newRegisterMessage("foo.bar")

	go func(t *testing.T) {
		eventBus, err := NewEventBus("localhost:7000")
		if err != nil {
			t.Error("Event bus creation failed", err)
		}
		defer eventBus.Close()

		if err = eventBus.send(sendMsg); err != nil {
			t.Error("Message sending failed", err)
		}
	}(t)

	conn, err := listener.Accept()
	if err != nil {
		t.Error("Accept() failed")
	}
	msg, err := receive(conn)
	if err != nil {
		t.Error("Bad response", err)
	}
	t.Log(msg)

	assert.Equal(t, sendMsg, msg)
}

type ExampleBody struct {
	Alpha string      `json:"alpha"`
	Beta  interface{} `json:"beta"`
}

func TestMessageSend(t *testing.T) {

	listener, err := net.Listen("tcp", "localhost:7000")
	if err != nil {
		t.Error("Starting server failed", err)
	}
	defer listener.Close()

	exampleBody := ExampleBody{
		Alpha: "hello",
		Beta: map[string]interface{}{
			"hello": "world",
		},
	}

	exampleBodyBytes, err := json.Marshal(&exampleBody)
	assert.NoError(t, err)

	headers := map[string]interface{}{
		"key": "value",
	}

	headersBytes, err := json.Marshal(&headers)
	assert.NoError(t, err)

	sendMsg := newSendMessage("foo.bar", "", headersBytes, exampleBodyBytes)

	go func(t *testing.T) {
		eventBus, err := NewEventBus("localhost:7000")
		if err != nil {
			t.Error("Event bus creation failed", err)
		}
		defer eventBus.Close()

		if err = eventBus.send(sendMsg); err != nil {
			t.Error("Message sending failed", err)
		}
	}(t)

	conn, err := listener.Accept()
	if err != nil {
		t.Error("Accept() failed")
	}
	msg, err := receive(conn)
	if err != nil {
		t.Error("Bad response", err)
	}
	t.Log(msg)

	assert.Equal(t, sendMsg, msg)
}
