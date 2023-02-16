package storage

import "app/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Post() PostRepoI
}

type UserRepoI interface {
	GetUsersList(search string) ([]models.User, error)
	UpdateUser(user models.User) (models.User, error)
	GetUserById(id string) (models.User, error)
	DeleteUserById(id string) (models.User, error)
}

type PostRepoI interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPostById(id string) (models.Post, error)
	GetAll() ([]models.Post, error)
	GetUserPosts(userId string) ([]models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
	DeletePost(id string) (models.Post, error)
}
