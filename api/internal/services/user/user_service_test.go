package user

import (
	"commerce/internal/shared/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	password := "hashed_password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepo := NewMockUserRepositoryI(ctl)
	mockRepo.EXPECT().GetByEmail("jon.doe@example.com").Return(&models.User{
		Base: models.Base{
			Id:          1,
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		Email:     "jon.doe@example.com",
		Password:  string(hashedPassword),
		FirstName: "Jon",
		LastName:  "Doe",
	}, nil)
	svc := NewUserService(mockRepo)
	user, err := svc.Authenticate("jon.doe@example.com", password)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "jon.doe@example.com", user.Email)
	assert.Equal(t, "Jon", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
}
