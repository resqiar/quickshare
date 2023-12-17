package services

import (
	"fmt"
	"quickshare/entities"
	"quickshare/repositories"
)

type UserService interface {
	RegisterUser(profile *entities.GooglePayload) (string, error)
	FindUserByEmail(email string) (*entities.User, error)
	FindUserByID(id string) (*entities.User, error)
}

type UserServiceImpl struct {
	UtilService UtilService
	Repository  repositories.UserRepo
}

func (service *UserServiceImpl) RegisterUser(profile *entities.GooglePayload) (string, error) {
	// format the given name from the provider
	formattedName := service.UtilService.FormatUsername(profile.GivenName)

	// concatenate formatted name with the nano id
	formattedName = fmt.Sprintf("%s_%s", formattedName, service.UtilService.GenerateRandomID(7))

	newUser := &entities.User{
		Username:   formattedName,
		Email:      profile.Email,
		PictureURL: profile.Picture,
	}

	result, err := service.Repository.CreateUser(newUser)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (service *UserServiceImpl) FindUserByEmail(email string) (*entities.User, error) {
	result, err := service.Repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *UserServiceImpl) FindUserByID(id string) (*entities.User, error) {
	result, err := service.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
