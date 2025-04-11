//go:build nocoverage
// +build nocoverage

package user

type User struct {
	Id    string
	Name  string
	Email string
}

// Repository defines the contract for user data operations.
type Repository interface {
	GetByID(id string) (*User, error)
	Save(user *User) error
}
