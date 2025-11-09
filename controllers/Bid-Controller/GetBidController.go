package bidcontroller

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

// GetBidData godoc
// @Summary      List applicants for a project
// @Description  Returns all bids submitted for the specified project
// @Tags         bids
// @Produce      json
// @Param        id   path      int  true  "Project ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /projects/{id}/applicants [get]
func GetBidData(c *gin.Context) {
	clientId := c.Param("id")
	client_id, err := strconv.Atoi(clientId)

	if err != nil {
		fmt.Println("The error in converting the client id is : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid project ID"})
		return
	}
	var applicants models.Applicants

	dbClient := lib.DBConnect()
	dbCollection := dbClient.Database("skillBackend").Collection("applicants")
	filter := bson.M{
		"client_id": client_id,
	}
	err = dbCollection.FindOne(context.TODO(), filter).Decode(&applicants)
	if err != nil {
		fmt.Println("The error in getting the docs is  : ", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "No applicants found for this project"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "true", "data": applicants})
}
