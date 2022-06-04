package entity

import "github.com/google/uuid"

type Conversation struct {
	Conversation_id uuid.UUID   // Primary Key
	User_id_list    []uuid.UUID // Foregn key
	Message_id_list []uuid.UUID // Foregn key
}
