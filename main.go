package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Sarthak-Java1124/go-SkillLink.git/lib"
	"github.com/Sarthak-Java1124/go-SkillLink.git/routes"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println("Starting server on port:", port)
	lib.DBConnect()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://w3ins-six.vercel.app"},
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

	router.Run(":" + port)

}
