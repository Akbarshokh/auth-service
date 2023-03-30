package rest

import (
	"auth_service/internal/errs"
	"auth_service/internal/models"
	"database/sql"
	"net/http"

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
			Return(ctx, nil, errs.Errf(errs.ErrValidation, err.Error()))
			return
		}

		//Verifying Refresh Token
		if err := verifyRefreshtoken(token.RefreshToken); err != nil {
			Return(ctx, nil, errs.New("Invalid Refresh Token"))
			return
		}
		//Parsing Client_id from refresh token
		refresh_token, _ := jwt.Parse(token.RefreshToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		claims, _ := refresh_token.Claims.(jwt.MapClaims)
		ClientID := claims["client_id"].(string)

		//Generating new Access and Refresh token
		access_token_str, refresh_token_str, err := generateTokens(ClientID)
		if err != nil {
			Return(ctx, nil, errs.New("Invalid Refresh Token"))
			return
		}

		//Updating tokens in db
		if err := updateTokensInDB(db, ClientID, access_token_str, refresh_token_str); err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrInternal, err.Error()))
			return
		}

		// Returning Tokens
		result := ReturnResult(ClientID, access_token_str, refresh_token_str)
		ctx.JSON(http.StatusOK, result)
	}
}

func verifyRefreshtoken(refresh_token string) error {
	_, err := jwt.Parse(refresh_token, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	return err
}

func updateTokensInDB(db *sql.DB, ClientID string, access_token_str string, refresh_token_str string) error {
	query := "UPDATE users SET access_token=$1, refresh_token=$2 WHERE client_id=$3"
	_, err := db.Exec(query, access_token_str, refresh_token_str, ClientID)
	if err != nil {
		return err
	}
	return nil
}

func ReturnResult(ClientID string, access_token_str string, refresh_token_str string) models.GetTokenRes{
	result := models.GetTokenRes{
		SignUpRes: models.SignUpRes{
			ClientID:     ClientID,
			AccessToken:  access_token_str,
			RefreshToken: refresh_token_str,
			Active:       true,
		},
	}
	return result
}
