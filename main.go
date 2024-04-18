package main

import (
	"fmt"
	"net/http"
	"user/userManagement"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	userManagement.Init()

	router := gin.Default()
	router.GET("/users", getUsers)         //to get all users
	router.GET("/users/:id", getUserByID)  // to get user by ID
	router.POST("/user", postUser)         // to create new user
	router.PATCH("/user/:id", updateUser)  // to updated existing user
	router.POST("/auth", authenticateUser) // authentication
	router.Run()

}

func getUsers(c *gin.Context) {
	fmt.Println("Main: getUsers!")
	c.IndentedJSON(http.StatusOK, userManagement.GetUsers(c))
}

func getUserByID(c *gin.Context) {
	fmt.Println("Main: getUserByID!")
	c.IndentedJSON(http.StatusOK, userManagement.GetUserByID(c))
}

func postUser(c *gin.Context) {
	fmt.Println("Main: postUser!")
	c.IndentedJSON(http.StatusCreated, userManagement.CreateUser(c))
}

func updateUser(c *gin.Context) {
	fmt.Println("Main: updateUser!")
	c.IndentedJSON(http.StatusCreated, userManagement.UpdateUser(c))
}

func authenticateUser(c *gin.Context) {
	fmt.Println("Main: authenticateUser!")
	c.IndentedJSON(http.StatusOK, userManagement.AuthenticateUser(c))
}
