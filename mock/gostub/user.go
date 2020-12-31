package gostub

type User struct {
	Name string
}

var GetUser = func(id int64) (*User, error) {
	return nil, nil
}

func Get(id int64) (*User, error) {
	return nil, nil
}

func (u *User) Get(id int64) (*User, error) {
	return nil, nil
}
