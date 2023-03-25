package rest

import (
	"database/sql"
	"jwt-go/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetToken(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token models.SignInRes
		if err := ctx.ShouldBindJSON(&token); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Verify Refresh Token
		_, err := jwt.Parse(token.RefreshToken, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid refresh token"})
			return
		}

		//Generate new Access Token
		access_token_expiration := time.Now().Add(time.Hour * 24 * 7)

		access_claims := jwt.MapClaims{
			"client_id": token.ClientID,
			"exp":       access_token_expiration.Unix(),
		}
		access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)
		access_token_str, err := access_token.SignedString([]byte("secret"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid refresh token"})
			return
		}
		//Update tokens in db
		query := "UPDATE users SET access_token=$1 WHERE refresh_token=$2"
		_, err = db.Exec(query, access_token_str, token.RefreshToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid refresh token"})
			return
		}
		// Return Tokens
		ctx.JSON(http.StatusOK, gin.H{
			"client_id":     token.ClientID,
			"access_token":  access_token_str,
			"refresh_token": token.RefreshToken,
		})
	}
}
