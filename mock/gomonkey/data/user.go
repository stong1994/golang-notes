package gomonkey

type User struct {
	Name string
}

func (u *User) Get(id int64) (string, error) {
	return u.Name, nil
}
