package rest

import (
	"auth_service/internal/errs"
	"auth_service/internal/models"
	"database/sql"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// type Response struct {
// 	Status    string      `json:"status"`
// 	ErrorCode int         `json:"error_code"`
// 	ErrorNote string      `json:"error_note"`
// 	Data      interface{} `json:"data"`
// }

// SignUp godoc
// @Router /sign-up [POST]
// @Summary Sign Up
// @Description API for Sign Up
// @Tags User Auth Service
// @Produce json
// @Param request body models.SignUpReq true "Client ID"
// @Success 201 {object} Response{data=models.SignUpRes}
// @Failure 400 {object} Response{data}
// @Failure 500 {object} Response{data}
func SignUp(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Parsing request body
		var signUpReq models.SignUpReq
		if err := ctx.ShouldBindJSON(&signUpReq); err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrValidation, err.Error()))
			return
		}
		//validating user input
		if err := validateUserInput(signUpReq, db); err != nil{
			Return(ctx, nil, errs.Errf(errs.ErrValidation, err.Error()))
			return
		}

		//generating hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpReq.Password), bcrypt.DefaultCost)
		if err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrInternal, err.Error()))
			return
		}

		//Generating access and refresh token
		access_token_str, refresh_token_str, err := generateTokens(signUpReq.ClientID)
		if err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrInternal, err.Error()))
			return
		}

		//Inserting into db
		err = insertUser(db, signUpReq, string(hashedPassword), access_token_str, refresh_token_str)
		if err != nil {
			Return(ctx, nil, errs.Errf(errs.ErrInternal, err.Error()))
			return
		}

		//Returning token to user
		signUpRes := models.SignUpRes{
			AccessToken:  access_token_str,
			RefreshToken: refresh_token_str,
			Active:       true,
			ClientID:     signUpReq.ClientID,
		}
		ctx.JSON(http.StatusOK, signUpRes)
	}
}

func validateUserInput(signUpReq models.SignUpReq, db *sql.DB) error{
	//Checking for empty fields
	if signUpReq.ClientID == "" || signUpReq.FirstName == "" ||
		signUpReq.LastName == "" || signUpReq.Email == "" || signUpReq.DeviceNum == "" || 
		signUpReq.Password == "" ||signUpReq.DeviceType == "" {
			return errs.New("Every field should not be empty")
	}
	//Checking for valid email address
	if !IsValidEmail(signUpReq.Email) {
		return errs.New("Invalid email")
	}
	//checking for unique email and client_id
	isUnique, err := IsUserUnique(db, signUpReq.Email, signUpReq.ClientID)
	if err != nil {
		return errs.Errf(errs.ErrInternal, err.Error())
	}
	if !isUnique {
		return errs.New("User with the same email or client_id already exists")
	}
	//checking for valid first_name and last_name (at least 2 chars)
	if len(signUpReq.FirstName) < 2 || len(signUpReq.LastName) < 2 {
		return errs.New("First Name and Last Name should be at least 2 charachters")
	}
	//checking password length (at least 8 characters, at least one uppercase letter, one lowercase letter, one digit, and one special character)
	if len(signUpReq.Password) < 8 {
		return errs.New("Password length should be at least 8 characters")
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range signUpReq.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper =true
		case unicode.IsLower(char):
			hasLower =true
		case unicode.IsDigit(char):
			hasDigit =true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return errs.New("Password should contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
	}
	return nil
}

func generateTokens(ClientID string) (string, string, error) {
	// Creating access token
	access_claims := jwt.MapClaims{
		"client_id": ClientID,
		"exp":       time.Now().Add(time.Hour * 24 * 7).Unix(), //expiration seven days
	}
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)
	access_token_str, err := access_token.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}

	// Creating refresh token
	refresh_claims := jwt.MapClaims{
		"client_id": ClientID,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(), //expiration one month
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)
	refresh_token_str, err := refresh_token.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}

	return access_token_str, refresh_token_str, nil
}

func insertUser(db *sql.DB, signUpReq models.SignUpReq, hashedPassword, accessToken string, refreshToken string) error {
	//Insert New User data into DB
	query := "INSERT INTO users (client_id, first_name, last_name, email, password, device_num, device_type, access_token, refresh_token, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	_, err := db.Exec(query, signUpReq.ClientID, signUpReq.FirstName, signUpReq.LastName, signUpReq.Email, hashedPassword, signUpReq.DeviceNum, signUpReq.DeviceType, accessToken, refreshToken, true)
	if err != nil {
		return err
	}

	return nil
}

func IsUserUnique(db *sql.DB, email string, ClientID string) (bool, error) {
	//checking for unique email and client_id
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1 OR client_id = $2", email, ClientID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func IsValidEmail(email string) bool {
	//Checking for valid email address
	if strings.Count(email, "@") != 1 {
		return false
	}
	return strings.Contains(email, "@hamkorbank.uz")
}

