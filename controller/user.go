package controller

import "app/models"

func (c *Controller) GetUsersListController(search string) ([]models.User, error) {
	users, err := c.store.User.GetUsersList(search)
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}

func (c *Controller) UpdateUserController(user models.User) (models.User, error) {
	user, err := c.store.User.UpdateUser(user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (c *Controller) GetUserByIdController(id string) (models.User, error) {
	user, err := c.store.User.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (c *Controller) DeleteUserByIdController(id string) (models.User, error) {
	user, err := c.store.User.DeleteUserById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
