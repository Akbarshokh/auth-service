package rest

import (
	"auth_service/internal/models"
	"auth_service/internal/errs"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// SignIn godoc
// @Router /sign-in [POST]
// @Summary Sign In using client_id, email, and access_token
// @Description API for Sign In
// @Tags User Auth Service
// @Produce json
// @Param request body models.SignInReq true "Client ID"
// @Success 201 {object} Response{data}
// @Failure 400 {object} Response{data}
// @Failure 500 {object} Response{data}
func SignIn(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signInReq models.SignInReq
		if err := ctx.ShouldBindJSON(&signInReq); err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrValidation, err.Error()))
			return
		}
		//Verifying is user exist in DB
		isUnique, err := IsUserUnique(db, signInReq.Email, signInReq.ClientID)
		if err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrValidation, err.Error()))
			return
		}
		if isUnique {
			Return(ctx, nil, errs.Errf(errs.ErrAuthorization, err.Error()))
			return
		}
		if !verifyAccessToken(signInReq.AccessToken) {
			Return(ctx, nil, errs.New("Invalid Access Token"))
			return
		}
		//Generating new Access Token
		newAccessToken, err := generateAccessToken(signInReq.ClientID)
		if err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrValidation, err.Error()))
			return
		}
		//Updating new access token in db
		err = updateAccessToken(db, signInReq.ClientID, newAccessToken)
		if err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrValidation, err.Error()))
			return
		}
		// Returning new access token
		ctx.JSON(http.StatusOK, gin.H{
			"access_token": newAccessToken,
		})
	}
}

func verifyAccessToken(token string) bool {
	// Parse the token with a secret key.
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	// Checking if the token is valid and has not expired.
	if err == nil && parsedToken.Valid {
		return true
	}

	return false
}

func generateAccessToken(ClientID string) (string, error) {
	// Creating access token
	access_claims := jwt.MapClaims{
		"client_id": ClientID,
		"exp":       time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)
	// sign the token with a secret key
	return token.SignedString([]byte("secret"))
}

func updateAccessToken(db *sql.DB, email string, accessToken string) error {
	query := "UPDATE users SET access_token = $1 WHERE email = $2"
	_, err := db.Exec(query, accessToken, email)
	return err
}

// if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(signUpReq.Password)); err != nil {
		// 	Return(ctx, nil, errs.New("Invalid email or password"))
		// 	return
		// }