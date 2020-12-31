package gomonkey

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"golang-learning/mock/monkey/data"
	"reflect"
	"testing"
)

// go test -gcflags=-l -v user_test.go user.go  -test.run TestUser_Get
func TestUser_Get(t *testing.T) {
	u := new(User)
	gomonkey.ApplyMethod(
		reflect.TypeOf(u),
		"Get",
		func(u *User, id int64) (string, error) {
			return "lee", nil
		},
	)
	msg, _ := u.Get(1)
	assert.Equal(t, "lee", msg)
}

/**
go test -gcflags=-l -v user_test.go data/user.go  -test.run TestUser_Get
报错：named files must all be in one directory; have ./ and data/
对于跨目录的文件如何关联测试？ TODO
*/
// 对于
func TestUser_Get2(t *testing.T) {
	u := new(data.User)
	gomonkey.ApplyMethod(
		reflect.TypeOf(u),
		"Get",
		func(u *data.User, id int64) (string, error) {
			return "lee", nil
		},
	)
	msg, _ := u.Get(1)
	assert.Equal(t, "lee", msg)
}
