package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charlygame/CatGameService/constants"
	"github.com/charlygame/CatGameService/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	err := userRepository.Get(id, &result)
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

func (s *UserService) Insert(document User) (interface{}, *utils.GameError) {
	id, err := userRepository.Create(document)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *UserService) Count(document interface{}) (int64, *utils.GameError) {
	count, err := userRepository.Count(document)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *UserService) WXLogin(code string) (*WXAuth, *utils.GameError) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", constants.WX_APP_ID, constants.WX_APP_SECRET, code)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return nil, &utils.GameError{StatusCode: 503, Err: err}
	}
	defer response.Body.Close()

	var result WXAuth = WXAuth{}

	json.NewDecoder(response.Body).Decode(&result)

	fmt.Printf("%+v\n", result)

	return &result, nil
}

func (s *UserService) WXRegister(user *User) (string, *utils.GameError) {
	id, err := s.Insert(*user)
	if err != nil {
		return "", &utils.GameError{StatusCode: 500, Err: err}
	}
	return id.(primitive.ObjectID).Hex(), nil
}

func (s *UserService) FindOne(query interface{}) (*User, *utils.GameError) {
	var result User
	err := userRepository.FindOne(query, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *UserService) GetRankList(skip int64, limit int64) ([]User, *utils.GameError) {
	var result []User

	projection := bson.D{{Key: "score", Value: 1}, {Key: "username", Value: 1}, {Key: "avatar", Value: 1}}
	order := bson.D{{Key: "score", Value: -1}}
	err := userRepository.List(nil, projection, skip, limit, order, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
