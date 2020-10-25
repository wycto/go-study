package math

type User struct {
	Name string
	age  string
	sex  string
}

func (User *User) SetName(value string) {
	User.Name = value
}

func (User *User) GetName() interface{} {
	return User.Name
}
