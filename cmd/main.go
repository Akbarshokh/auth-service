package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"jwt-go/internal/rest"
	"jwt-go/api/docs"
	// "jwt-go/internal/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dbsql"
	password = "gosql"
	dbname   = "jwt_go"
)

func main() {
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	//Connect to db
	psql := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psql)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	router := gin.Default()

	router.POST("/sign-up", rest.SignUp(db))
	router.POST("/check-token", rest.CheckToken(db))
	router.POST("/get-token", rest.GetToken(db))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := router.Run("172.25.0.32:8080"); err != nil {
		log.Fatal(err)
	}
}
