package structs

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Email string `json:"email"`
}
