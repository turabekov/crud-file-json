package controller

import "app/storage"

type Controller struct {
	store *storage.Store
}

// Controller Constructor
func NewController(store *storage.Store) *Controller {
	return &Controller{
		store: store,
	}
}
