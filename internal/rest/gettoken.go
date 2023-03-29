package rest

import (
	"database/sql"
	"auth_service/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GetToken godoc
// @Router /get-token [POST]
// @Summary Checking token with Refresh Token
// @Description This endpoint verifies token is active or not and generates new access token
// @Tags User Auth Service
// @Produce json
// @Param request body models.GetTokenReq true "Refresh Token"
// @Success 201 {object} Response{data=models.GetTokenRes}
// @Failure 400 {object} Response{data}
// @Failure 500 {object} Response{data}
func GetToken(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token models.GetTokenReq
		//Parsing request body
		if err := ctx.ShouldBindJSON(&token); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Verifying Refresh Token
		_, err := jwt.Parse(token.RefreshToken, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid refresh token"})
			return
		}
		var response models.GetTokenRes
		//Generating new Access Token
		access_token_expiration := time.Now().Add(time.Hour * 24 * 7)

		access_claims := jwt.MapClaims{
			"client_id": response.ClientID,
			"exp":       access_token_expiration.Unix(),
		}
		access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)
		access_token_str, err := access_token.SignedString([]byte("secret"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid refresh token"})
			return
		}

		// Creating refresh token
		refresh_claims := jwt.MapClaims{
			"client_id": response.ClientID,
			"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
		}
		refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)
		refresh_token_str, err := refresh_token.SignedString([]byte("secret"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid refresh token"})
			return
		}

		//Updating tokens in db
		query := "UPDATE users SET access_token=$1 WHERE refresh_token=$2"
		_, err = db.Exec(query, access_token_str, refresh_token_str)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid refresh token"})
			return
		}
		
		result := models.GetTokenRes{
			SignUpRes: models.SignUpRes{
				ClientID:     "client_id",
				AccessToken:  access_token_str,
				RefreshToken: refresh_token_str,
				Active:       true,
			},
		}
		// Returning Tokens
		ctx.JSON(http.StatusOK, result)
	}
}
