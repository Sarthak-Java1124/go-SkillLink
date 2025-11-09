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

// PostBids godoc
// @Summary      Submit a bid for a project
// @Description  Adds a bid to the applicants collection for the specified project
// @Tags         bids
// @Accept       json
// @Produce      json
// @Param        id    path      int           true  "Project ID"
// @Param        bid   body      models.Bids   true  "Bid payload"
// @Success      201   {object}  map[string]string
// @Success      202   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /projects/{id}/apply [post]
func PostBids(c *gin.Context) {
	clientId := c.Param("id")
	client_id, err := strconv.Atoi(clientId)
	if err != nil {
		fmt.Println("The error in converting the client id is : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid project ID"})
		return
	}
	var bidsData models.Bids

	if err := c.ShouldBind(&bidsData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot bind the bid json"})
		return
	}

	dbClient := lib.DBConnect()
	dbCollection := dbClient.Database("skillBackend").Collection("applicants")
	filter := bson.M{
		"client_id": client_id,
	}
	update := bson.M{
		"$push": bson.M{
			"applicants": bidsData,
		},
	}
	result, err := dbCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("The error in updating the collection is :", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "There is an error in updating the collection"})
		return
	}
	if result.MatchedCount == 0 {
		applicant := models.Applicants{
			ClientId:       client_id,
			ApplicantArray: []models.Bids{bidsData},
		}
		_, insertErr := dbCollection.InsertOne(context.Background(), applicant)
		if insertErr != nil {
			fmt.Println("The error in inserting the new applicant document is :", insertErr)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to create applicant record"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Successfully created applicant record"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Succesfully inserted the bid"})

}
