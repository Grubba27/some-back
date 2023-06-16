package auth

// auth is a package for authentication and authorization
// consider it as a User Management Service
// a User service, the flow should go always user -> auth

import (
	"app/api/v1/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type registerUserRequest struct {
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
//	@Param			body	body		registerUserRequest	true	"body"
//	@Success		200		{string}	id
//	@Failure		400		{string}	error
//	@Router			/auth/register [post]
func registerUser(g *gin.Context) {
	var r registerUserRequest

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
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(10 * time.Minute),
	})

	g.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

type loginUserReturn struct {
	Email         string `json:"email" binding:"required,email"`
	PublicAddress string `json:"publicAddress"`
	Token         string `json:"token"`
}

// LoginUser godoc
//
//	@Summary	 Login
//	@Schemes
//	@Description	Login user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	registerUserRequest	true	"body"
//	@Success		200		{object}	loginUserReturn
//	@Failure		400		{string}	error
//	@Router			/auth/login [post]
func login(g *gin.Context) {
	var r registerUserRequest

	if err := g.ShouldBindJSON(&r); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while binding JSON"})
		return
	}

	user, err := user.FindByEmail(r.Email)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while finding user"})
		return
	}
	
	hashedPassword, err := hashPassword(r.Password)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while hashing password"})
		return
	}

	if checkPassword(hashedPassword, user.Password) {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password", "description": "Invalid password"})
		return
	}

	token, err := createJWTToken(user.ID)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while creating JWT token"})
		return
	}

	http.SetCookie(g.Writer, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(10 * time.Minute),
	})
	result := loginUserReturn{
		Email:         user.Email,
		PublicAddress: user.PublicAddress,
		Token:         token,
	}
	g.JSON(http.StatusOK, result)
}

// here we create all routes for auth
func InitModule(g *gin.RouterGroup) {
	g.POST("/register", registerUser)
	g.POST("/login", login)
}
