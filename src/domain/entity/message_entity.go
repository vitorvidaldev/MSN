package entity

import "github.com/google/uuid"

type Message struct {
	Message_id      uuid.UUID // Primary key
	Sender_id       uuid.UUID // Foregn key
	Conversation_id uuid.UUID // Foregn key
	Message_body    string
}
