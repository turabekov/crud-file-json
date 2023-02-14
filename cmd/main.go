package main

import (
	"fmt"
	"log"

	"app/config"
	"app/controller"
	"app/models"
	"app/storage"
)

func main() {

	cfg := config.Load()

	store, err := storage.NewFileJSON(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, store)

	// Get list of user
	res, err := c.GetUsersListController("be")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	// Update user
	userForUpdate := models.User{
		Id:        "32a803b5-3f1c-495e-8ce7-4649e4cbe3b1",
		FirstName: "Khumoyun",
		LastName:  "Shaxzod",
		Gender:    "Male",
		Address: models.Address{
			Street: "A.Ibragimov 1-1",
			City:   "Tashkent",
		},
		Friends: []models.Friend{
			{
				Id:          "7ba900aa-75e8-4f17-9089-845446ff1515",
				Email:       "Darrell97@hotmail.com",
				PhoneNumber: "578-164-9681",
			},
			{
				Id:          "7ba900aa-7qe8-4f17-9089-845446ff1515",
				Email:       "khumoyun@gmail.com",
				PhoneNumber: "777-332-3333",
			},
		},
		CardNumber: "342465392896553",
		Birthday:   "2002-06-12",
		Profession: "Backend developer",
	}

	user1, err := c.UpdateUserController(userForUpdate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user1)

	// Get user by id

	user2, err := c.GetUserByIdController("c57aa672-902f-44c8-af9d-dfa02f62541a")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Get By Id", user2)

	// Delete user by id

	user3, err := c.DeleteUserByIdController("c57aa672-902f-44c8-af9d-dfa02f62541a")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted user", user3)
}
