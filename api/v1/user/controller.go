package user

import (
	"app/lib/cripto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserDTO struct {
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
		Email:         user.Email,
		PublicAddress: user.PublicAddress,
	}
	g.JSON(200, gin.H{"user": response})
}

type PublicAddressDTO struct {
	Message			string `json:"message"`
	Signature		string `json:"signature"`
}

type PublicAddressResponse struct {
	PublicAddress	string `json:"publicAddress"`
}
// UpdateUserPublicAdress godoc
//
//	@Summary	Update current public adress user
//	@Security	jwt
//	@Schemes
//	@Description	Update current user public adress
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			body	body		PublicAddressDTO	true	"body"
//	@Success		200		{object}	PublicAddressResponse
//	@Failure		400		{string}	error
//	@Router			/user/profile [put]
func updateProfile(g *gin.Context) {
	userID := g.MustGet("userID").(uint)
	var r PublicAddressDTO

	if err := g.ShouldBindJSON(&r); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "description": "Error while binding JSON"})
		return
	}

	user, err := FindByID(userID)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error(), "description": "Error while getting user"})
		return	
	}
	user.PublicAddress = cripto.GeneratePublicAddress(r.Message, r.Signature)
	
	if user.PublicAddress == "" {
		g.JSON(400, gin.H{"error": "Invalid signature", "description": "Error while updating user"})
		return
	}

	result, err := UpdateUser(user)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error(), "description": "Error while updating user"})
		return
	}

	g.JSON(200, gin.H{"user": PublicAddressResponse{
		PublicAddress: result.PublicAddress,
	}})
}

// GetUserByAddress godoc
//
//	@Summary	Get user by address
//	@Security	jwt
//	@Schemes
//	@Description	Get user by address
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			publicAddress	path		string	true	"publicAddress"
//	@Success		200				{object}	UserDTO
//	@Failure		400				{string}	error
//	@Router			/user/{publicAddress} [get]
func getProfileByAddress(g *gin.Context) {
	publicAddress := strings.ReplaceAll(g.Param("publicAddress"), "/", "")
	user, err := FindByPublicAddress(publicAddress)

	if err != nil {
		g.JSON(400, gin.H{"error": err.Error(), "description": "Error while getting user"})
		return
	}

	response := UserDTO{
		Email:         user.Email,
		PublicAddress: user.PublicAddress,
	}
	g.JSON(200, gin.H{"user": response})
}

// is an Authenticated route
// context.get("userID") is the  current user id
func InitModule(g *gin.RouterGroup) {
	g.GET("/profile", getMe)
	g.GET("/:publicAddress", getProfileByAddress)
	g.PUT("/profile", updateProfile)
}
