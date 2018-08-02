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

// Message A Vert.x event bus message that can be mapped to JSON.
type Message struct {
	Type         string `json:"type"`
	ReplyAddress string `json:"replyAddress,omitempty"`
	Address      string `json:"address"`
	Headers      []byte `json:"headers"`
	Body         []byte `json:"body"`
}

// IsError This method returns true if the Type field has value 'err'.
func (msg *Message) IsError() bool {
	return msg.Type == "err"
}

func newSendMessage(address string, replyAddress string, headers, body []byte) *Message {
	return &Message{
		Type:         "send",
		Address:      address,
		ReplyAddress: replyAddress,
		Headers:      headers,
		Body:         body,
	}
}

func newPublishMessage(address string, headers, body []byte) *Message {
	return &Message{
		Type:    "publish",
		Address: address,
		Headers: headers,
		Body:    body,
	}
}

func newRegisterMessage(address string) *Message {
	return &Message{
		Type:    "register",
		Address: address,
	}
}

func newUnregisterMessage(address string) *Message {
	return &Message{
		Type:    "unregister",
		Address: address,

		Headers: nil,
		Body:    nil,
	}
}
