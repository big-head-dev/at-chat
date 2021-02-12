package main

// StatusMessage for non-user messages
type StatusMessage struct {
	Type    string `json:"type"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

func newStatusMessage(message string) StatusMessage {
	return StatusMessage{
		Type:    "status",
		Time:    getTimeNow("3:04PM"),
		Message: message,
	}
}

// ChatMessage for basic text messages
type ChatMessage struct {
	Type     string `json:"type"`
	Time     string `json:"time"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func newChatMessage(username string, message string) ChatMessage {
	return ChatMessage{
		Type:     "chat",
		Time:     getTimeNow("3:04PM"),
		Username: username,
		Message:  message,
	}
}
