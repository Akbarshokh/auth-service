package rest

import (
	"database/sql"
	"jwt-go/internal/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CheckToken godoc
// @Router /check-token [POST]
// @Summary Checking token with Access Token
// @Description This endpoint verifies token is active or not
// @Tags Product Service
// @Produce json
// @Param request body user.User true "Client ID"
// @Success 201 {object} Response{data}
// @Failure 400 {object} Response{data}
// @Failure 500 {object} Response{data}
func CheckToken(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token user.User
		//Parse request body
		if err := ctx.ShouldBindJSON(&token); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Verify Acces Token
		_, err := jwt.Parse(token.AccessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"active": false})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"active": true})
	}
}
