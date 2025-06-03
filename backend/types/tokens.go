package types

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TokStr    string             `bson:"tokStr"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone"`
	Role      Role               `bson:"role"`
	CreatedAt time.Time          `bson:"createdAt"`
	Expires   time.Time          `bson:"expires"`
}

var ErrSignupTokenNotExist = errors.New("Signup token does not exist")

type SignupTokenStore interface {
	CreateSignupToken(Phone, Email string, Role Role, Expiry time.Duration) (*SignupToken, error)
	GetSignupToken(TokStr string) (*SignupToken, error)
}
