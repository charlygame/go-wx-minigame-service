package user

import (
	"github.com/charlygame/CatGameService/repository"
)

const (
	timeout = 5
)

type UserRepository struct {
	repository.MongoRepository
}

func NewUserRepository() UserRepository {
	r := UserRepository{}
	r.Collection = "users"
	return r
}
