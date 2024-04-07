package structs

import "time"

type Role string

const (
	Rider   Role = "Rider"
	admin   Role = "Admin"
	support Role = "Support"
	Driver  Role = "Driver"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PublicUser struct {
	ID                string    `json:"id" bson:"_id,omitempty"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	Gender            string    `json:"gender"`
	Role              string    `json:"role"`
	ProfilePictureUrl string    `json:"profilePictureUrl"`
	CreatedAt         time.Time `json:"createdAt"`
	IsActive          bool      `json:"isActive"`
}

type PrivateUser struct {
	ID                string    `json:"id" bson:"_id,omitempty"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	Gender            string    `json:"gender"`
	Role              string    `json:"role"`
	ProfilePictureUrl string    `json:"profilePictureUrl"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	IsEmailVerified   bool      `json:"isEmailVerified"`
	IsActive          bool      `json:"isActive"`
	FcmToken          string    `json:"fcmToken"`
}
