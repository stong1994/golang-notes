package user

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mocks/mock"
	"testing"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	id := int64(1)
	mockMale := mock.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(nil),
	)

	user := NewUser(mockMale)
	err := user.GetUserInfo(id)
	assert.NoError(t, err)
}
