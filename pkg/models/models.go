package models

// Message is the message struct that is sent between clients and the server.
// Message is the content of the message
// Document is the name of the document that the message belongs to
type Message struct {
	Message  string `json:"message"`
	Document string `json:"document"`
}

// CreateMessage creates a new message struct.
func CreateMessage(message string, document string) Message {
	return Message{
		Message:  message,
		Document: document,
	}
}
