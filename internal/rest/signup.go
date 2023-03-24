package rest

import (
	"database/sql"
	"jwt-go/internal/models"
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
// func SignIn(db *sql.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		//Parse request body
// 		var user user.User
// 		if err := ctx.ShouldBindJSON(&user); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		// create JWT tokens
// 		access_token_expiration := time.Now().Add(time.Hour * 24 * 7)   //  7 days
// 		refresh_token_expiration := time.Now().Add(time.Hour * 24 * 30) // 30 days

// 		access_claims := jwt.MapClaims{
// 			"client_id": user.ClientID,
// 			"exp":       access_token_expiration.Unix(), // expiration time for access token
// 		}
// 		access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)
// 		access_token_str, err := access_token.SignedString([]byte("secret"))
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		refresh_claims := jwt.MapClaims{
// 			"client_id": user.ClientID,
// 			"exp":       refresh_token_expiration.Unix(), // expiration time for refresh token
// 		}
// 		refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)
// 		refresh_token_str, err := refresh_token.SignedString([]byte("secret"))
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Insert into db
// 		query := "INSERT INTO users (client_id, first_name, last_name, email, device_num, device_type, access_token, refresh_token, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
// 		var userID int
// 		errr := db.QueryRow(query, user.ClientID, user.FirstName, user.LastName, user.Email, user.DeviceNum, user.DeviceType, access_token_str, refresh_token_str, true).Scan(&userID)
// 		if errr != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		//Return tokens to user
// 		ctx.JSON(http.StatusCreated, gin.H{
// 			"access_token":  access_token_str,
// 			"active":        true,
// 			"client_id":     user.ClientID,
// 			"refresh_token": refresh_token_str,
// 		})
// 	}
// }

func SignIn(db *sql.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//Parse request body
		var signInReq models.SignInReq
		if err := ctx.ShouldBindJSON(&signInReq); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Generate 	access and refresh token
		access_token_str, refresh_token_str, err := generateTokens(signInReq.ClientID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}

		//Insert into db
		err = insertUser(db, signInReq, access_token_str, refresh_token_str)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error})
			return
		}

		//Return token to user
		signInRes := models.SignInRes{
			AccessToken : access_token_str,
			RefreshToken : refresh_token_str,
			Active : true,
			ClientID: signInReq.ClientID,
		}
		ctx.JSON(http.StatusOK, signInRes)
	}
}

func generateTokens(ClientID string) (string, string, error) {
	// Create access token
	access_claims := jwt.MapClaims{
		"client_id": ClientID,
		"exp":       time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)
	access_token_str, err := access_token.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}

	// Create refresh token
	refresh_claims := jwt.MapClaims{
		"client_id": ClientID,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)
	refresh_token_str, err := refresh_token.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}

	return access_token_str, refresh_token_str, nil
}

func insertUser(db *sql.DB, signInReq models.SignInReq, accessToken string, refreshToken string) error {
	query := "INSERT INTO users (client_id, first_name, last_name, email, device_num, device_type, access_token, refresh_token, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := db.Exec(query, signInReq.ClientID, signInReq.FirstName, signInReq.LastName, signInReq.Email, signInReq.DeviceNum, signInReq.DeviceType, accessToken, refreshToken, true)
	if err != nil {
		return err
	}

	return nil
}