package main

import "encoding/json"

// StatusMessage for non-user messages
type StatusMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func newStatusMessage(message string) StatusMessage {
	return StatusMessage{
		Type:    "status",
		Message: message,
	}
}

//json.Marshal helper function
func (c StatusMessage) toJSON() ([]byte, error) {
	return json.Marshal(c)
}

// ChatMessage for basic text messages
type ChatMessage struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func newChatMessage(username string, message string) ChatMessage {
	return ChatMessage{
		Type:     "chat",
		Username: username,
		Message:  message,
	}
}

//json.Marshal helper function
func (c ChatMessage) toJSON() ([]byte, error) {
	return json.Marshal(c)
}

// ChatMessages for sending a collection of text messages
type ChatMessages struct {
	Type     string        `json:"type"`
	Messages []ChatMessage `json:"messages"`
}

//json.Marshal helper function
func (c ChatMessages) toJSON() ([]byte, error) {
	return json.Marshal(c)
}
