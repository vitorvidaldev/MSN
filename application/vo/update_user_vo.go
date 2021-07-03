package vo

type UpdateUserVO struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
}
