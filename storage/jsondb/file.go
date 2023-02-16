package jsondb

import (
	"app/config"
	"app/storage"
	"os"
)

type Store struct {
	user *userRepo
	post *postRepo
}

func NewFileJSON(cfg *config.Config) (*Store, error) {

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	postFile, err := os.Open(cfg.Path + cfg.PostFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		user: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
		post: NewPostRepo(cfg.Path+cfg.PostFileName, postFile),
	}, nil

}

func (s *Store) CloseDB() {
	s.user.file.Close()
	s.post.file.Close()
}

func (s *Store) User() storage.UserRepoI {
	return s.user
}

func (s *Store) Post() storage.PostRepoI {
	return s.post
}
