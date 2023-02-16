package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type postRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewPostRepo(fileName string, file *os.File) *postRepo {
	return &postRepo{
		fileName: fileName,
		file:     file,
	}
}

func readPostsFile() ([]models.Post, error) {
	posts := []models.Post{}
	// read json file
	data, err := ioutil.ReadFile("./data/posts.json")
	if err != nil {
		return nil, err
	}
	// json parse
	err = json.Unmarshal(data, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Create===================================================
func (p *postRepo) CreatePost(post models.Post) (models.Post, error) {
	var posts []*models.Post

	err := json.NewDecoder(p.file).Decode(&posts)
	if err != nil {
		return models.Post{}, err
	}

	id := uuid.NewString()

	newPost := &models.Post{
		Id:          id,
		Title:       post.Title,
		Description: post.Description,
		UserId:      post.UserId,
	}

	posts = append(posts, newPost)

	body, err := json.MarshalIndent(posts, "", "   ")

	if err != nil {
		return models.Post{}, err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return models.Post{}, err
	}

	return *newPost, nil
}

// Get by id post===================================================
func (p *postRepo) GetPostById(id string) (models.Post, error) {
	posts, err := readPostsFile()
	if err != nil {
		return models.Post{}, err
	}

	for _, v := range posts {
		if v.Id == id {
			return models.Post{
				Id:          v.Id,
				Title:       v.Title,
				Description: v.Description,
				UserId:      v.UserId,
			}, nil
		}
	}

	return models.Post{}, errors.New("post not found")
}

// // Get All posts===================================================
func (p *postRepo) GetAll() ([]models.Post, error) {
	posts, err := readPostsFile()
	if err != nil {
		return []models.Post{}, err
	}

	return posts, nil
}

// // get user posts =============================================================

func (p *postRepo) GetUserPosts(userId string) ([]models.Post, error) {
	posts, err := readPostsFile()
	if err != nil {
		return []models.Post{}, err
	}
	usersPost := []models.Post{}
	for _, v := range posts {
		if userId == v.UserId {
			usersPost = append(usersPost, v)
		}
	}

	return usersPost, nil
}

// // Update post ===============================================================
func (p *postRepo) UpdatePost(post models.Post) (models.Post, error) {
	posts, err := readPostsFile()
	if err != nil {
		return models.Post{}, err
	}

	updatedPost := models.Post{}
	for i, v := range posts {
		if v.Id == post.Id {
			posts[i].Title = post.Title
			posts[i].Description = post.Description
			posts[i].UserId = post.UserId
			updatedPost = posts[i]
		}
	}

	if len(updatedPost.Id) <= 0 {
		return models.Post{}, errors.New("Post not found")
	}

	body, err := json.MarshalIndent(posts, "", "   ")

	if err != nil {
		return models.Post{}, err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return models.Post{}, err
	}

	return updatedPost, nil

}

// // Delete Post ==============================================================
func (p *postRepo) DeletePost(id string) (models.Post, error) {
	posts, err := readPostsFile()
	if err != nil {
		return models.Post{}, err
	}

	deletedPost := models.Post{}
	for i, v := range posts {
		if v.Id == id {
			deletedPost = posts[i]
			posts = append(posts[:i], posts[i+1:]...)
		}
	}

	if len(deletedPost.Id) <= 0 {
		return models.Post{}, errors.New("Post not found")
	}

	body, err := json.MarshalIndent(posts, "", "   ")

	if err != nil {
		return models.Post{}, err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return models.Post{}, err
	}

	return deletedPost, nil
}
