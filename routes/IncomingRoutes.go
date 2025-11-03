package routes

import (
	"github.com/Sarthak-Java1124/go-SkillLink.git/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users/:id")
	incomingRoutes.PATCH("/users/:id")
}

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/auth/register", controllers.AuthControllerSignUp)
	incomingRoutes.POST("/auth/login")
}

func ProjectRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/projects")
	incomingRoutes.POST("/projects/:id/apply")
	incomingRoutes.POST("/projects/:id/accept/:freelancerId")
	incomingRoutes.GET("/projects/:id/applicants")
}

func ContractRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/contracts/:id")
	incomingRoutes.POST("/contracts/:id/submit-work")
	incomingRoutes.POST("/contracts/:id/approve")
	incomingRoutes.POST("/contract/:id/dispute")
}
