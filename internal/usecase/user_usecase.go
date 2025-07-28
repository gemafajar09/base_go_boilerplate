package usecase

import (
	"errors"
	"go-project/internal/auth"
	"go-project/internal/domain"
	"go-project/internal/repository"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

// Register user baru
func (uc *UserUsecase) Register(user domain.User) (domain.User, error) {
	// Hash password sebelum simpan
	user.Password = auth.HashPassword(user.Password)
	return uc.userRepository.CreateUser(user)
}

// Login dengan email & password, return JWT token kalau berhasil
func (uc *UserUsecase) Login(email, password string) (string, error) {
	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Verify password bcrypt
	if err := auth.CheckPasswordHash(password, user.Password); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.Id, user.Email, string(user.Role))
	if err != nil {
		return "", err
	}

	return token, nil
}
