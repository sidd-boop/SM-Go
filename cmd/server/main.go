package main

import (
	"SM-Go/internal/db"
	"SM-Go/internal/middleware"
	"SM-Go/internal/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"SM-Go/internal/auth"
	"SM-Go/internal/thought"
)


func main(){
	if err:=godotenv.Load(); err!=nil{
		log.Fatal("Error in .env file loading.")
	}
	r:= gin.New()

	r.Use(middleware.LoggingMiddleware())
	db.Connect()
	
	//Migration error check
	if err:=auth.Migrate(db.DB);err!=nil{
		log.Fatal("Auth migration failed:", err)
	}

	if err:=thought.Migrate(db.DB);err!=nil{
		log.Fatal("Thought migration failed", err)
	}
	router.RoutesHandler(r)

	err := r.Run(":"+os.Getenv("APP_PORT"))
	if err!=nil{
		log.Fatal("Failed to start server: ", err)
	}
}