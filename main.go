package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/Sarthak-Java1124/go-SkillLink.git/docs"
	"github.com/Sarthak-Java1124/go-SkillLink.git/lib"
	"github.com/Sarthak-Java1124/go-SkillLink.git/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Your API Name
// @version         1.0
// @description     Short description of your API.
// @termsOfService  https://example.com/terms
// @contact.name    API Support
// @contact.email   dev@example.com
// @license.name    MIT
// @host            localhost:8080
// @BasePath        /
// @schemes         http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println("Starting server on port:", port)
	lib.DBConnect()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	routes.ProjectRoutes(router)
	routes.ContractRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + port)

}
