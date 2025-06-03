package auth

import (
	"time"

	"github.com/jeremiafourie/cogniflight-cloud/backend/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FakeUserStore struct {
	Users map[string]types.User
}

func (s FakeUserStore) GetUserByEmail(email string) (*types.User, error) {
	user, ok := s.Users[email]

	if !ok {
		return nil, types.ErrUserNotExist
	} else {
		return &user, nil
	}
}

type FakeSessionStore struct {
	Sessions     map[string]types.Session
	CreateCalled bool
	UserID       primitive.ObjectID
	Role         types.Role
	SessID       string
}

func (s *FakeSessionStore) CreateSession(UserID primitive.ObjectID, Role types.Role) (*types.Session, error) {
	s.CreateCalled = true
	s.UserID = UserID
	s.Role = Role
	sessID, err := GenerateToken()
	if err != nil {
		return nil, err
	}

	if s.Sessions == nil {
		s.Sessions = map[string]types.Session{}
	}
	s.Sessions[sessID] = types.Session{
		ID:        primitive.NewObjectID(),
		SessID:    sessID,
		UserID:    UserID,
		Role:      Role,
		CreatedAt: time.Now(),
	}

	s.SessID = sessID
	return &types.Session{
		ID:        primitive.NewObjectID(),
		UserID:    UserID,
		Role:      Role,
		CreatedAt: time.Now(),
		SessID:    sessID,
	}, nil
}

func (s FakeSessionStore) GetSession(SessID string) (*types.Session, error) {
	session, ok := s.Sessions[SessID]

	if !ok {
		return nil, types.ErrSessionNotExist
	} else {
		return &session, nil
	}
}
