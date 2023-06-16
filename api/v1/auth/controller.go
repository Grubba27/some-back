package auth

import (
	"app/api/v1/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterUser godoc
//
//	@Summary	Register user
//	@Schemes
//	@Description	Register new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		RegisterUserRequest	true	"body"
//	@Success		200		{string}	id
//	@Failure		400		{string}	error
//	@Router			/auth/register [post]
func registerUser(g *gin.Context) {
	var r RegisterUserRequest

	if err := g.ShouldBindJSON(&r); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while binding JSON"})
		return
	}

	hashedPassword, err := hashPassword(r.Password)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while hashing password"})
		return
	}

	user, err := user.Create(r.Email, hashedPassword)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while creating user"})
		return
	}

	token, err := createJWTToken(user.ID)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while creating JWT token"})
		return
	}

	http.SetCookie(g.Writer, &http.Cookie{
		Name: 	 "token",
		Value: 	 token,
		Expires: time.Now().Add(10 * time.Minute),
	})
	
	g.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// here we create all routes for auth module
func InitModule(g *gin.RouterGroup) {
	g.POST("/register", registerUser)
}
