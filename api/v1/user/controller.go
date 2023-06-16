package user

import "github.com/gin-gonic/gin"

type userResponse struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	PublicAddress string `json:"publicAddress"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

// GetUser godoc
//
//			@Summary	Get current user
//	   @Security jwt
//			@Schemes
//			@Description	Get current user
//			@Tags			user
//			@Accept			json
//			@Produce		json
//			@Success		200		{object}	userResponse
//			@Failure		400		{string}	error
//			@Router			/user/profile [get]
func getMe(g *gin.Context) {
	userID := g.MustGet("userID").(uint)
	user, err := FindByID(userID)

	if err != nil {
		g.JSON(400, gin.H{"error": err.Error(), "description": "Error while getting user"})
		return
	}
	response := userResponse{
		ID:            user.ID,
		Email:         user.Email,
		PublicAddress: user.PublicAddress,
		CreatedAt:     user.CreatedAt.String(),
		UpdatedAt:     user.UpdatedAt.String(),
	}
	g.JSON(200, gin.H{"user": response})
}

// is an Authenticated route
// context.get("userID") is the  current user id

func InitModule(g *gin.RouterGroup) {
	g.GET("profile", getMe)
}
