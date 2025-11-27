package main

import (
	"database/sql"
	"log"

	_ "github.com/momomo0206/rest-api-in-go-gin/docs"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"github.com/momomo0206/rest-api-in-go-gin/internal/database"
	"github.com/momomo0206/rest-api-in-go-gin/internal/env"
)

// @title Go Gin Rest API
// @version 1.0
// @description A Rest API in Go using Gin framework
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-123456"),
		models:    models,
	}
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
