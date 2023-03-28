package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sethvargo/go-envconfig"

	"jwt-go/api/docs"
	"jwt-go/internal/config"
	"jwt-go/internal/cors"
	"jwt-go/internal/rest"

	// "jwt-go/internal/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	_ = godotenv.Load()
	var cfg config.Config
	if err := envconfig.ProcessWith(context.TODO(), &cfg, envconfig.OsLookuper()); 
	err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Host = cfg.ServerIP + cfg.HTTPPort
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	//Connecting to db
	db, err := sql.Open("postgres", cfg.PostgresURI())
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

	router.Use(cors.CORSMiddleware())

	router.POST("/sign-up", rest.SignUp(db))
	router.POST("/sign-in", rest.SignIn(db))
	router.POST("/check-token", rest.CheckToken(db))
	router.POST("/get-token", rest.GetToken(db))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := router.Run("172.25.0.32:8080"); err != nil {
		log.Fatal(err)
	}
}
