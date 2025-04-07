package domain

type User struct {
	UserId      string
	UserName    string
	PhoneNumber string
	Gender      string
}

type UserRepository interface {
	Create(user *User) error
}
