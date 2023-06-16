package main

import (
	auth "app/api/v1/auth"
	user "app/api/v1/user"
	"app/db"
	docs "app/docs"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin" // swagger embed files
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//	@BasePath	/api/v1

// @Summary	Health Check
// @Schemes
// @Description	do ping
// @Tags			health
// @Accept			json
// @Produce		json
// @Success		200	{string}	ok
// @Router			/health [get]
func HealthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @securityDefinitions.apikey	jwt
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load()
	if err != nil && os.Getenv("RAILWAY") != "true" {
		log.Fatal("Error loading .env file")
	}
	{
		db.InitDB()
		db.GetDB().AutoMigrate(&user.User{})
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	v1.GET("/health", HealthCheck)
	{
		auth.InitModule(v1.Group("/auth"))
	}
	{
		userGroup := v1.Group("user")
		userGroup.Use(auth.IsAuthorized())
		user.InitModule(userGroup)
	}
	if os.Getenv("PROD") == "false" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	log.Println("Server is running on port " + os.Getenv("PORT"))
	port := ":" + os.Getenv("PORT")
	r.Run(port)
}
