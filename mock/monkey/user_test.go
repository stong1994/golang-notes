package monkey

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// go test -gcflags=-l -v user_test.go user.go  -test.run TestUser_Get
func TestUser_Get(t *testing.T) {
	u := new(User)
	monkey.PatchInstanceMethod(
		reflect.TypeOf(u),
		"Get",
		func(u *User, id int64) (string, error) {
			return "lee", nil
		},
	)
	msg, _ := u.Get(1)
	assert.Equal(t, "lee", msg)
}
