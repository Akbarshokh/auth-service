package rest

import (
	"auth_service/internal/models"
	"auth_service/internal/errs"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CheckToken godoc
// @Router /check-token [POST]
// @Summary Checking token with Access Token
// @Description This endpoint verifies token is active or not
// @Tags User Auth Service
// @Produce json
// @Param request body models.CheckTokenReq true "Access Token"
// @Success 201 {object} Response{data=models.CheckTokenRes}
// @Failure 400 {object} Response{data}
// @Failure 500 {object} Response{data}
func CheckToken(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token models.CheckTokenReq
		//Parsing request body
		if err := ctx.ShouldBindJSON(&token); err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrValidation, err.Error()))
			return
		}
		//Verifying Acces Token
		_, err := jwt.Parse(token.AccessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"active": false})
			return
		}
		result := models.CheckTokenRes{
			Active: true,
		}
		ctx.JSON(http.StatusOK, result)
	}
}
