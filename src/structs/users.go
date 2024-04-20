package structs

import (
	"errors"
	"time"
)

type Role string

const (
	User    Role = "User"
	Admin   Role = "Admin"
	Support Role = "Support"
)


type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rq *UserRequest) ValidateUserRequest() error {
	if rq.Email == "" || len(rq.Email) < 5 {
		return errors.New("invalid Email")
	}
	if rq.Password == "" || len(rq.Password) < 6 {
		return errors.New("invalid Password")
	}
	return nil
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
	AccessToken       string    `json:"accessToken"`
}
