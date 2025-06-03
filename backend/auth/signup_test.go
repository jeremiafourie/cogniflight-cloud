package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jeremiafourie/cogniflight-cloud/backend/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FakeSignupTokenStore struct {
	Tokens       map[string]types.SignupToken
	CreateCalled bool
	Created      *types.SignupToken
}

func (s *FakeSignupTokenStore) CreateSignupToken(Phone, Email string, Role types.Role, Expiry time.Duration) (*types.SignupToken, error) {
	tokStr, err := GenerateToken()
	s.CreateCalled = true

	if err != nil {
		return nil, err
	}

	tok := types.SignupToken{
		ID:        primitive.NewObjectID(),
		TokStr:    tokStr,
		Email:     Email,
		Phone:     Phone,
		Role:      Role,
		CreatedAt: time.Now(),
	}
	tok.Expires = tok.CreatedAt.Add(Expiry)

	if s.Tokens == nil {
		s.Tokens = map[string]types.SignupToken{}
	}

	s.Tokens[tokStr] = tok
	s.Created = &tok

	return &tok, nil
}

func (s *FakeSignupTokenStore) GetSignupToken(TokStr string) (*types.SignupToken, error) {
	tok, ok := s.Tokens[TokStr]

	if !ok {
		return nil, types.ErrSignupTokenNotExist
	} else {
		return &tok, nil
	}
}

func TestCreateSignupToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	tokenStore := FakeSignupTokenStore{}
	r.POST("/create-signup-token", CreateSignupToken(&tokenStore))

	t.Run("Invalid request body is 400", func(t *testing.T) {
		badBodies := []string{
			"",
			`{}`,
			`{"email": "example@gmail.com"}`,
			`{"phone": "271738749839"}`,
			`{"role": "pilot"}`,
			`{"em`,
		}

		for i, body := range badBodies {
			t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
				w := FakeRequest(t, r, "POST", body, "/create-signup-token", nil)

				t.Logf("Request body: %q", body)
				if w.Result().StatusCode != 400 {
					t.Errorf("Wrong StatusCode: want %d got %d", 400, w.Result().StatusCode)
				}
			})
		}
	})
	t.Run("Valid request succeeds", func(t *testing.T) {
		goodBodies := []string{
			`{"email": "example@gmail.com", "phone": "271738749839", "role": "pilot"}`,
			`{"email": "example@gmail.com", "role": "pilot"}`,
			`{"phone": "271738749839", "role": "pilot"}`,
		}

		for i, body := range goodBodies {
			t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
				w := FakeRequest(t, r, "POST", body, "/create-signup-token", nil)

				t.Logf("Request body: %q", body)
				if w.Result().StatusCode != 201 {
					t.Errorf("Wrong StatusCode: want %d got %d", 201, w.Result().StatusCode)
				}
				if !tokenStore.CreateCalled {
					t.Error("Expected create to be called on SignupTokenStore")
				} else if tokenStore.Created == nil {
					t.Error("Expected tokenStore to have the created token")
				}

			})
			tokenStore = FakeSignupTokenStore{}
		}
	})
}
