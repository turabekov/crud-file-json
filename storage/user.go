package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"app/models"
)

type userRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewUserRepo(fileName string, file *os.File) *userRepo {
	return &userRepo{
		fileName: fileName,
		file:     file,
	}
}

func readFile() ([]models.User, error) {
	users := []models.User{}
	// read json file
	data, err := ioutil.ReadFile("./data/users.json")
	if err != nil {
		return nil, err
	}
	// json parse
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// get all user
func (u *userRepo) GetUsersList(search string) ([]models.User, error) {
	users, err := readFile()
	if err != nil {
		return []models.User{}, err
	}
	// logics of search functionality
	res := []models.User{}
	for _, v := range users {
		fullName := strings.ToLower(v.FirstName + v.LastName)
		if strings.Contains(fullName, strings.ToLower(strings.Replace(search, " ", "", -1))) {
			res = append(res, v)
		}
	}

	// if user not found
	if len(res) <= 0 {
		return res, errors.New("user doesn't exist")
	}

	return res, nil
}

// update user
func (u *userRepo) UpdateUser(user models.User) (models.User, error) {
	users, err := readFile()
	if err != nil {
		return models.User{}, err
	}

	updatedUser := models.User{}
	for i, v := range users {
		if v.Id == user.Id {
			users[i].FirstName = user.FirstName
			users[i].LastName = user.LastName
			users[i].Gender = user.Gender
			users[i].Address = user.Address
			users[i].Friends = user.Friends
			users[i].CardNumber = user.CardNumber
			users[i].Birthday = user.Birthday
			users[i].Profession = user.Profession

			updatedUser = users[i]
		}
	}

	if len(updatedUser.Id) <= 0 {
		return models.User{}, errors.New("user not found")
	}

	body, err := json.MarshalIndent(users, "", "   ")
	if err != nil {
		return models.User{}, err
	}

	err = ioutil.WriteFile("./data/users.json", body, os.ModePerm)
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

// get user by id
func (u *userRepo) GetUserById(id string) (models.User, error) {
	users, err := readFile()
	if err != nil {
		return models.User{}, err
	}

	for _, v := range users {
		if v.Id == id {
			return v, nil
		}
	}

	return models.User{}, errors.New("user not found")
}

// delete user by id
func (u *userRepo) DeleteUserById(id string) (models.User, error) {
	users, err := readFile()
	if err != nil {
		return models.User{}, err
	}

	deletedUser := models.User{}
	for i, v := range users {
		if v.Id == id {
			deletedUser = v
			users = append(users[:i], users[i+1:]...)
		}
	}

	if len(deletedUser.Id) <= 0 {
		return models.User{}, errors.New("user not found")
	}

	body, err := json.MarshalIndent(users, "", "   ")
	if err != nil {
		return models.User{}, err
	}

	err = ioutil.WriteFile("./data/users.json", body, os.ModePerm)
	if err != nil {
		return models.User{}, err
	}

	return deletedUser, nil
}
