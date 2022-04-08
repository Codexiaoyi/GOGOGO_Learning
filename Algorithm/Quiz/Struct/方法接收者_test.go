package struct1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type IUser interface {
	Modify()
	ModifyWithPointer()
}

type User struct {
	email string
	name  string
}

func (u User) Modify() {
	u.email = "111"
	u.name = "1111"
}

func (u *User) ModifyWithPointer() {
	u.email = "222"
	u.name = "2222"
}

func TestModified(t *testing.T) {
	user := User{email: "000", name: "0000"}
	userPoint := &User{email: "000", name: "0000"}

	user.Modify()
	assert.Equal(t, "000", user.email)
	userPoint.Modify()
	assert.Equal(t, "000", userPoint.email)

	user.ModifyWithPointer()
	assert.Equal(t, "222", user.email)
	userPoint.ModifyWithPointer()
	assert.Equal(t, "222", userPoint.email)
}

func TestInterface(t *testing.T) {
	// var user IUser
	// var userPoint IUser
	// user = User{email: "000", name: "0000"} //这里报错
	// userPoint = &User{email: "000", name: "0000"}
}
