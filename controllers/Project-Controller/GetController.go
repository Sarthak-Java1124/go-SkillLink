package projectcontroller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sarthak-Java1124/go-SkillLink.git/lib"
	"github.com/Sarthak-Java1124/go-SkillLink.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetFormData godoc
// @Summary      Get project by client ID
// @Description  Retrieves a project associated with the provided client ID
// @Tags         projects
// @Produce      json
// @Param        id   path      int  true  "Client ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Router       /project/{id} [get]
func GetFormData(c *gin.Context) {
	clientId := c.Param("id")
	clientId_int, err := strconv.Atoi(clientId)
	if err != nil {
		fmt.Println("The error in conversion of string is : ", clientId_int)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid client ID"})
		return
	}
	fmt.Println("The id we are getting from the params is  :", clientId)
	filters := bson.M{"client_id": clientId_int}
	var GetData models.ProjectModel
	dbClient := lib.DBConnect()
	collection := dbClient.Database("skillBackend").Collection("projects")
	err = collection.FindOne(context.Background(), filters).Decode(&GetData)
	if err != nil {
		fmt.Println("The error in getting the docs from the db is : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error in searching file from the db"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "true", "data": GetData})
	}

}
