package user

import "github.com/charlygame/CatGameService/utils"

type UserService struct {
}

var userRepository UserRepository

func NewUserService() UserService {
	userRepository = NewUserRepository()
	s := UserService{}
	return s
}

func (s *UserService) Get(id string) (User, *utils.GameError) {
	var result User
	err := userRepository.Get(id, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *UserService) Update(id string, document User) (User, *utils.GameError) {
	err := userRepository.Update(id, document)
	if err != nil {
		return document, err
	}
	return document, nil
}

func (s *UserService) Insert(document User) *utils.GameError {
	_, err := userRepository.Create(document)
	if err != nil {
		return err
	}
	return nil
}
