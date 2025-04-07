package user

import "errors"

type InMemoryRepo struct {
	users map[string]*User
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{users: make(map[string]*User)}
}

func (r *InMemoryRepo) GetByID(id string) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryRepo) Save(user *User) error {
	r.users[user.Id] = user
	return nil
}
