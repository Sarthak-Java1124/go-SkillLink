package projectcontroller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Sarthak-Java1124/go-SkillLink.git/lib"
	"github.com/Sarthak-Java1124/go-SkillLink.git/models"
	"github.com/gin-gonic/gin"
)

// PostProjectController godoc
// @Summary      Create a new project
// @Description  Stores a project document in the database
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        project  body      models.ProjectModel  true  "Project payload"
// @Success      201      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Router       /projects [post]
func PostProjectController(c *gin.Context) {
	var UserForm models.ProjectModel

	if err := c.ShouldBindJSON(&UserForm); err != nil {
		fmt.Println("The error in binding the json is : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error in binding the json for form"})
		return
	}

	dbClient := lib.DBConnect()
	dbCollection := dbClient.Database("skillBackend").Collection("projects")
	_, err := dbCollection.InsertOne(context.TODO(), UserForm)
	if err != nil {
		fmt.Println("The error in saving the file in the db is : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error saving file in the db"})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "Form submitted successfully"})
	}

}
