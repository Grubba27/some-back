package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserDTO struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	PublicAddress string `json:"publicAddress"`
}

// GetUser godoc
//
//	@Summary	Get current user
//	@Security	jwt
//	@Schemes
//	@Description	Get current user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	UserDTO
//	@Failure		400	{string}	error
//	@Router			/user/profile [get]
func getMe(g *gin.Context) {
	userID := g.MustGet("userID").(uint)
	user, err := FindByID(userID)

	if err != nil {
		g.JSON(400, gin.H{"error": err.Error(), "description": "Error while getting user"})
		return
	}

	response := UserDTO{
		ID:            user.ID,
		Email:         user.Email,
		PublicAddress: user.PublicAddress,
	}
	g.JSON(200, gin.H{"user": response})
}

func updateProfile(g *gin.Context) {
	userID := g.MustGet("userID").(uint)
	var r UserDTO

	if err := g.ShouldBindJSON(&r); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while binding JSON"})
		return
	}

	user := User{
		Email:         r.Email,
		PublicAddress: r.PublicAddress,
	}
	user.ID = userID
	result, err := UpdateUser(user)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error(), "description": "Error while updating user"})
		return
	}

	g.JSON(200, gin.H{"user": result})
}

// is an Authenticated route
// context.get("userID") is the  current user id
func InitModule(g *gin.RouterGroup) {
	g.GET("profile", getMe)
}
