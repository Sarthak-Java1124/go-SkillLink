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

	"go.mongodb.org/mongo-driver/bson"
)

func AuthControllerSignUp(c *gin.Context) {
	var body models.UserModel

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("The error in binding the json is :", err)
		c.JSON(http.StatusBadGateway, gin.H{"message": "There was an error in binding json"})
	}

	dbClient := lib.DBConnect()
	userCollection := dbClient.Database("skilllink").Collection("users")
	isPresent, err := userCollection.CountDocuments(context.TODO(), bson.M{"email": body.Email})
	if err != nil {
		fmt.Println("The error in counting docs is : ", err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"message": "There was an error in counting documents",
		})
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
	}

}

type LoginStruct struct {
	Email    *string
	Password *string
}

func AuthControllerLogin(c *gin.Context) {
	var body LoginStruct

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("The error in binding login json is : ", err)
		c.JSON(http.StatusForbidden, gin.H{"message": "Error in binding login json"})

	}
	fmt.Println("The binded json is : ", *&body)
	dbClient := lib.DBConnect()
	userInstance := dbClient.Database("skilllink").Collection("users")
	var UserBody models.UserModel
	err := userInstance.FindOne(context.TODO(), bson.M{"email": body.Email}).Decode(&UserBody)
	if err != nil {
		fmt.Println("The error in finding the user by login email is : ", err)
		return

	}
	comparePassword := utils.VerifyHashPassword(*body.Password, *UserBody.Password)
	if !comparePassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The password didn't match"})
	}
}
