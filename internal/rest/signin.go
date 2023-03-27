package rest

import (
	"database/sql"
	"jwt-go/internal/models"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func SignIn(db *sql.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var signInReq models.SignInReq
		if err := ctx.ShouldBindJSON(&signInReq);
		err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Verify is user exist in DB
		
	}
}