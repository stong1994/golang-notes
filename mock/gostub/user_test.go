package gostub

import (
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 模拟一个函数变量
func TestGetUser(t *testing.T) {
	stubs := gostub.StubFunc(&GetUser, &User{Name: "lee"}, nil)
	defer stubs.Reset()
	user, err := GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, "lee", user.Name)
}

// 模拟一个函数
func TestGet(t *testing.T) {
	var get = Get
	stubs := gostub.StubFunc(&get, &User{Name: "lee"}, nil)
	defer stubs.Reset()
	user, err := get(1)
	assert.NoError(t, err)
	assert.Equal(t, "lee", user.Name)
}
