package user

type UserService struct {
}

var userRepository UserRepository

func NewUserService() UserService {
	userRepository = NewUserRepository()
	s := UserService{}
	return s
}

func (s *UserService) Get(userID string) (User, error) {
	// user, err := userRepository.Get(userID)
	return nil, nil
}
