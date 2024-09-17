package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"log"
)

type UserService struct {
}

func (s *UserService) FindUserByEmail(email string) (*domain.User, error) {
	// Implement the logic to find a user by email

	return nil, nil
}

func (s *UserService) SingUp(input dto.UserRegister) (string, error) {
	// Implement the logic to create a user
	log.Println(input)

	return "", nil
}

func (s *UserService) Login(input any) (string, error) {
	// Implement the logic to create a user

	return "", nil
}

func (s *UserService) GetVerificationCode(user domain.User) (int, error) {
	// Implement the logic to create a user

	return 0, nil
}

func (s *UserService) VerifyCode(id uint, code int) error {
	// Implement the logic to create a user

	return nil
}

func (s *UserService) CreateProfile(id uint, input any) error {
	// Implement the logic to create a user

	return nil
}

func (s *UserService) GetProfile(id uint) (*domain.User, error) {
	// Implement the logic to create a user

	return nil, nil
}

func (s *UserService) UpdateProfile(id uint, input any) error {
	// Implement the logic to create a user

	return nil
}

func (s *UserService) BecomeSeller(id uint, input any) (string, error) {
	// Implement the logic to create a user

	return "", nil
}

func (s *UserService) FindCart(id uint) ([]interface{}, error) {
	// Implement the logic to create a user

	return nil, nil
}

func (s *UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	// Implement the logic to create a user

	return nil, nil
}

func (s *UserService) CreateOrder(u domain.User) (int, error) {
	// Implement the logic to create a user

	return 0, nil
}

func (s *UserService) GetOrders(U domain.User) ([]interface{}, error) {
	// Implement the logic to create a user

	return nil, nil
}

func (s *UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	// Implement the logic to create a user

	return nil, nil
}
