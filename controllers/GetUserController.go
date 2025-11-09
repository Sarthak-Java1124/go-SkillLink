package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Sarthak-Java1124/go-SkillLink.git/lib"
	"github.com/Sarthak-Java1124/go-SkillLink.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetUser godoc
// @Summary      Get user by ID
// @Description  Fetches a user document using its MongoDB ObjectID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ObjectID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("The id we get from the params is : ", id)
	
	// Use bson.ObjectIDFromHex from v2 instead of primitive.ObjectIDFromHex from v1
	mongoId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("The error in parsing mongoId is :", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "false",
			"message": "Invalid ID format",
		})
		return
	}
	fmt.Println("The type of mongoId is : ", fmt.Sprintf("%T", mongoId))
	fmt.Println("The mongo id is : ", mongoId)

	var user models.UserModel
	dbClient := lib.DBConnect()
	userInstance := dbClient.Database("skillBackend").Collection("users")

	// Simple query with v2 compatible ObjectID
	err = userInstance.FindOne(context.TODO(), bson.M{"_id": mongoId}).Decode(&user)
	if err != nil {
		fmt.Println("The error coming from the database : ", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "false",
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "true",
		"data":   user,
	})
}
