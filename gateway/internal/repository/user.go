package repository

type User struct {
	Login    string
	Password string
}

type UserRepository interface {
	GetUserByLogin(login string) *User
}

type UserRepositoryMock struct {
	users []*User
}

func NewUserRepositoryMock() *UserRepositoryMock {
	var users = make([]*User, 0)

	users = append(users, &User{"user1", "password"})
	users = append(users, &User{"user2", "password"})

	return &UserRepositoryMock{users: users}
}

func (u UserRepositoryMock) GetUserByLogin(login string) *User {
	var user *User

	for _, user = range u.users {
		if user.Login == login {
			return user
		}
	}

	return nil
}
