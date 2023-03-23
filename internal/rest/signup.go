package rest

import (
	"database/sql"
	"jwt-go/internal/models/user"
	"net/http"
	"time"
	// "crypto/ecdsa"
	// "crypto/elliptic"
	// "crypto/rand"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Response struct {
	Status    string
	ErrorCode int
	ErrorNote string
	Data      interface{}
}



// SignIn godoc
// @Router /sign [POST]
// @Summary Sign In
// @Description API for Sign In
// @Tags Product Service
// @Produce json
// @Param request body models.SignInReq true "Client ID"
// @Success 201 {object} Response{data=models.SignInRes}
// @Failure 400 {object} Response{data}
// @Failure 500 {object} Response{data}
func Signin(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}


// SignIn godoc
// @Router /sign [POST]
// @Summary Sign In
// @Description API for Sign In
// @Tags Product Service
// @Produce json
// @Param request body user.User true "Client ID"
// @Success 201 {object} Response{data=}
// @Failure 400 {object} Response{data}
// @Failure 500 {object} Response{data}
func Signup(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
// SignIn godoc
// @Router /sign [POST]
// @Summary Sign In
// @Description API for Sign In
// @Tags Product Service
// @Produce json
// @Param request body user.User true "Client ID"
// @Success 201 {object} Response{data}
// @Failure 400 {object} Response{data}
// @Failure 500 {object} Response{data}
func SignIn(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Parse request body
		var user user.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// create JWT tokens
		access_token_expiration := time.Now().Add(time.Hour * 24 * 7)   //  7 days
		refresh_token_expiration := time.Now().Add(time.Hour * 24 * 30) // 30 days

		access_claims := jwt.MapClaims{
			"client_id": user.ClientID,
			"exp":       access_token_expiration.Unix(), // expiration time for access token
		}
		access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)
		access_token_str, err := access_token.SignedString([]byte("secret"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		refresh_claims := jwt.MapClaims{
			"client_id": user.ClientID,
			"exp":       refresh_token_expiration.Unix(), // expiration time for refresh token
		}
		refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)
		refresh_token_str, err := refresh_token.SignedString([]byte("secret"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Insert into db
		query := "INSERT INTO users (client_id, first_name, last_name, email, device_num, device_type, access_token, refresh_token, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
		var userID int
		errr := db.QueryRow(query, user.ClientID, user.FirstName, user.LastName, user.Email, user.DeviceNum, user.DeviceType, access_token_str, refresh_token_str, true).Scan(&userID)
		if errr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//Insert tokens into db
		// query = "INSERT INTO tokens (user_id, access_token, refresh_token, active) VALUES ($1, $2, $3, true)"
		// _, err = db.Exec(query, userID, access_token, refresh_token)
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }

		//Return tokens to user
		ctx.JSON(http.StatusCreated, gin.H{
			"access_token":  access_token_str,
			"active":        true,
			"client_id":     user.ClientID,
			"refresh_token": refresh_token_str,
		})
	}
}
