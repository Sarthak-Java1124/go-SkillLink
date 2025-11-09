package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Sarthak-Java1124/go-SkillLink.git/lib"
	"github.com/Sarthak-Java1124/go-SkillLink.git/models"
	"github.com/Sarthak-Java1124/go-SkillLink.git/utils"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// AuthControllerSignUp godoc
// @Summary      Register a new user
// @Description  Creates a user account with the provided payload
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserModel  true  "User registration payload"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      417   {object}  map[string]string
// @Failure      502   {object}  map[string]string
// @Router       /auth/register [post]
func AuthControllerSignUp(c *gin.Context) {
	var body models.UserModel

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("The error in binding the json is :", err)
		c.JSON(http.StatusBadGateway, gin.H{"message": "There was an error in binding json"})
		return
	}

	dbClient := lib.DBConnect()
	userCollection := dbClient.Database("skillBackend").Collection("users")
	isPresent, err := userCollection.CountDocuments(context.TODO(), bson.M{"email": body.Email})
	if err != nil {
		fmt.Println("The error in counting docs is : ", err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"message": "There was an error in counting documents",
		})
		return
	}
	if isPresent > 0 {
		c.JSON(http.StatusBadGateway, gin.H{"error": "User already Exists"})
		return
	}

	hashedPassword, err := utils.HashThePassword(*body.Password)
	if err != nil {
		log.Fatal("The error in hashing the password is : ", err)
	}
	body.Password = &hashedPassword
	_, err = userCollection.InsertOne(context.TODO(), &body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Cannot save in the DB"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

type LoginStruct struct {
	Email    *string
	Password *string
}

// AuthControllerLogin godoc
// @Summary      Authenticate a user
// @Description  Validates user credentials and returns a session cookie
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginStruct  true  "User login credentials"
// @Success      202          {object}  map[string]string
// @Failure      400          {object}  map[string]string
// @Failure      403          {object}  map[string]string
// @Failure      424          {object}  map[string]string
// @Router       /auth/login [post]
func AuthControllerLogin(c *gin.Context) {
	var body LoginStruct

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("The error in binding login json is : ", err)
		c.JSON(http.StatusForbidden, gin.H{"message": "Error in binding login json"})
		return
	}
	dbClient := lib.DBConnect()
	userInstance := dbClient.Database("skillBackend").Collection("users")

	var UserBody models.UserModel
	err := userInstance.FindOne(context.TODO(), bson.M{"email": body.Email}).Decode(&UserBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No User with these credentials found"})
		fmt.Println("The error in finding the user by login email is : ", err)
		return
	}

	comparePassword := utils.VerifyHashPassword(*body.Password, *UserBody.Password)
	if !comparePassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The password didn't match"})

	} else {
		token, err := utils.GenerateJwt(UserBody.Name, UserBody.Email)
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{"message": "Error in parsing the jwt"})
		}
		c.SetCookie("token", token, 3600, "/", "localhost", true, true)
		c.JSON(http.StatusAccepted, gin.H{"message": "You are successfully logged in"})
	}

}
