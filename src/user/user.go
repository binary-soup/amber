package user

type User struct {
	Username     string
	PasswordHash []byte
}

func NewUser(username string, hash []byte) *User {
	return &User{
		Username:     username,
		PasswordHash: hash,
	}
}
