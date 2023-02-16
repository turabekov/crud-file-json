package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsondb"
	"fmt"
	"log"
)

func main() {

	cfg := config.Load()

	jsondb, err := jsondb.NewFileJSON(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, jsondb)

	// User======================================================================================================================================================
	// Get list of user
	// res, err := c.GetUsersListController("be")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

	// // Update user
	// userForUpdate := models.User{
	// 	Id:        "32a803b5-3f1c-495e-8ce7-4649e4cbe3b1",
	// 	FirstName: "Turabekov",
	// 	LastName:  "Khumoyun",
	// 	Gender:    "Male",
	// 	Address: models.Address{
	// 		Street: "A.Ibragimov 1-1",
	// 		City:   "Tashkent",
	// 	},
	// 	Friends: []models.Friend{
	// 		{
	// 			Id:          "7ba900aa-75e8-4f17-9089-845446ff1515",
	// 			Email:       "Darrell97@hotmail.com",
	// 			PhoneNumber: "578-164-9681",
	// 		},
	// 		{
	// 			Id:          "7ba900aa-7qe8-4f17-9089-845446ff1515",
	// 			Email:       "khumoyun@gmail.com",
	// 			PhoneNumber: "777-332-3333",
	// 		},
	// 	},
	// 	CardNumber: "342465392896553",
	// 	Birthday:   "2002-06-12",
	// 	Profession: "Backend developer",
	// }

	// user1, err := c.UpdateUserController(userForUpdate)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(user1)

	// // Get user by id

	// user2, err := c.GetUserByIdController("c57aa672-902f-44c8-af9d-dfa02f62541a")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Get By Id", user2)

	// Delete user by id

	// user3, err := c.DeleteUserByIdController("c57aa672-902f-44c8-af9d-dfa02f62541a")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Deleted user", user3)

	// Posts======================================================================================================================================================

	// Create
	// post, err := c.CreatePost(models.Post{
	// 	Title:       "sadasturkey holidays in Havai",
	// 	Description: "Hello dear",
	// 	UserId:      "ea7276d9-0931-4d8d-b8a9-e6d6c51e7584",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(post)

	// Get by Id
	// res, err := c.GetPostById("e76a16sd1f-c219-4bef-a637-8c65968e0081")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Get By id:", res)

	// Get All
	// posts, err := c.GetAllPosts()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Get All:", posts)

	// Get user posts

	posts, err := c.GetUserPosts("32a803b5-3f1c-495e-8ce7-4649e4cbe3b1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User posts:", posts)
	// Update post
	// p, e := c.UpdatePost(models.Post{
	// 	Id:          "e76a161f-c219-4bef-a637-8c65968e0081",
	// 	Title:       "Barcelona vs MUN Utd",
	// 	Description: "Good game",
	// 	UserId:      "32a803b5-3f1c-495e-8ce7-4649e4cbe3b1",
	// })
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// fmt.Println(p)

	// Deleted post
	// p, e := c.DeletePost("3315ebc5-935a-4fb7-82b2-4f021c15049a")

	// if e != nil {
	// 	log.Fatal(e)
	// }

	// fmt.Println(p)

}
