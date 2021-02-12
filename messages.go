package main

// ChatMessage for basic text messages
type ChatMessage struct {
	Type     string   `json:"type"`
	Time     string   `json:"time"`
	Username string   `json:"username,omitempty"`
	Message  string   `json:"message"`
	Users    []string `json:"users,omitempty"`
}

func newChatMessage(username string, message string) ChatMessage {
	return ChatMessage{
		Type:     "chat",
		Time:     getTimeNow("3:04PM"),
		Username: username,
		Message:  message,
	}
}

func newStatusMessage(message string, users []string) ChatMessage {
	return ChatMessage{
		Type:    "status",
		Time:    getTimeNow("3:04PM"),
		Message: message,
		Users:   users,
	}
}
