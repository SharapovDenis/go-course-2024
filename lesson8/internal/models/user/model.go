package user

type User struct {
	ID    int64
	Name  string
	Email string
}

func New() User {
	return User{}
}
