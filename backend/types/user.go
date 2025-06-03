package types

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	RolePilot    Role = "pilot"
	RoleATC      Role = "atc"
	RoleSysAdmin Role = "sysadmin"
)

type PilotInfo struct {
	FaceEmbeddings [][]float64    `bson:"faceEmbeddings"`
	Baseline       map[string]any `bson:"baseline"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	Phone        string             `bson:"phone"`
	Pwd          string             `bson:"pwd"`
	Role         Role               `bson:"role"`
	ProfileImage primitive.ObjectID `bson:"profileImage,omitempty"`
	PilotInfo    *PilotInfo         `bson:"pilotInfo,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt"`
}

var ErrUserNotExist = errors.New("User does not exist")

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(User User) (*User, error)
}
