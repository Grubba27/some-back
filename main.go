package main

import (
	auth "app/api/v1/auth"
	docs "app/docs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//	@BasePath	/api/v1

//	@Summary	Health Check
//	@Schemes
//	@Description	do ping
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	ok
//	@Router			/health [get]
func HealthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	v1.GET("/health", HealthCheck)
	{
		auth.InitModule(v1.Group("/auth"))
	}
	if os.Getenv("ENV") != "production" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	r.Run(":8080")
}
