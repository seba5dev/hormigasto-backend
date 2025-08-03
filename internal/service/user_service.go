package service

import (
	"errors"

	"github.com/seba5dev/hormigasto-backend/internal/repository"
	"github.com/seba5dev/hormigasto-backend/models"
)

type UserService interface {
	Register(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

// Example: Aquí podrías hashear la contraseña, validar datos, etc.
func (s *userService) Register(user *models.User) error {
	if user.Email == "" || user.PasswordHash == "" {
		return errors.New("email and password required")
	}
	// TODO: Hashear la contraseña aquí si no viene hasheada
	return s.repo.Create(user)
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) FindAll() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) Update(user *models.User) error {
	// Puedes agregar validaciones antes de actualizar
	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
