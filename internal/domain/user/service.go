package user

import "github.com/google/uuid"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetUser(id string) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) CreateUser(name, email string) (*User, error) {
	user := &User{
		Id:    uuid.NewString(),
		Name:  name,
		Email: email,
	}

	if err := s.repo.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}
