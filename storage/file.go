package storage

import "os"

type Store struct {
	User *userRepo
}

func NewFileJSON() (*Store, error) {

	userFile, err := os.Open("./data/users.json")
	if err != nil {
		return nil, err
	}

	return &Store{
		User: NewUserRepo("./data/users.json", userFile),
	}, nil

}
