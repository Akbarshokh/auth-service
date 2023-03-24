package rest

import (
	"database/sql"
	"jwt-go/internal/models"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Response struct {
	Status    string
	ErrorCode int
	ErrorNote string
	Data      interface{}
}


// SignUp godoc
// @Router /sign-up [POST]
// @Summary Sign Up
// @Description API for Sign Up
// @Tags Product Service
// @Produce json
// @Param request body models.SignInReq true "Client ID"
// @Success 201 {object} Response{data=models.SignInRes}
// @Failure 400 {object} Response{data}
// @Failure 500 {object} Response{data}
func SignUp(db *sql.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//Parse request body
		var signInReq models.SignInReq
		if err := ctx.ShouldBindJSON(&signInReq); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Verify user
		isUnique, err := isUserUnique(db, signInReq.Email, signInReq.ClientID)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        if !isUnique {
            ctx.JSON(http.StatusConflict, gin.H{"error": "User with the same email or client_id already exists"})
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

func isUserUnique(db *sql.DB, email string, ClientID string) (bool, error) {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1 OR client_id = $2", email, ClientID).Scan(&count)
    if err != nil {
        return false, err
    }
    return count == 0, nil
}
