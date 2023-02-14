package storage

import (
	"app/config"
	"os"
)

type Store struct {
	User *userRepo
}

func NewFileJSON(cfg *config.Config) (*Store, error) {

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		User: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
	}, nil

}
