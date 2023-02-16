package controller

import (
	"app/models"
	"fmt"
)

func (c *Controller) CreatePost(post models.Post) (models.Post, error) {
	userId := post.UserId

	_, err := c.store.User().GetUserById(userId)
	if err != nil {
		return models.Post{}, err
	}

	post, err = c.store.Post().CreatePost(post)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (c *Controller) GetPostById(id string) (models.ResponseGetByIdPost, error) {
	post, err := c.store.Post().GetPostById(id)
	if err != nil {
		return models.ResponseGetByIdPost{}, err
	}

	userid := post.UserId
	user, err := c.store.User().GetUserById(userid)
	if err != nil {
		return models.ResponseGetByIdPost{}, err
	}

	response := &models.ResponseGetByIdPost{
		Id:          post.Id,
		Title:       post.Title,
		Description: post.Description,
		User:        user,
	}

	return *response, nil

}

func (c *Controller) GetAllPosts() ([]models.ResponseGetByIdPost, error) {
	posts, err := c.store.Post().GetAll()
	if err != nil {
		return []models.ResponseGetByIdPost{}, err
	}

	res := []models.ResponseGetByIdPost{}
	for _, v := range posts {
		userId := v.UserId
		fmt.Println(userId)
		user, err := c.store.User().GetUserById(userId)
		if err != nil {
			return []models.ResponseGetByIdPost{}, err
		}

		res = append(res, models.ResponseGetByIdPost{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			User:        user,
		})
	}

	return res, nil
}

func (c *Controller) GetUserPosts(userId string) ([]models.Post, error) {
	_, err := c.store.User().GetUserById(userId)
	if err != nil {
		return []models.Post{}, err
	}

	posts, err := c.store.Post().GetUserPosts(userId)
	if err != nil {
		return []models.Post{}, err
	}

	return posts, nil
}

func (c *Controller) UpdatePost(post models.Post) (models.Post, error) {
	post, err := c.store.Post().UpdatePost(post)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (c *Controller) DeletePost(id string) (models.Post, error) {
	post, err := c.store.Post().DeletePost(id)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}
