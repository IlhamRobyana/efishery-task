package main

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("can't load .env : %v", err))
	}

	port := os.Getenv("APP_PORT")
	dsn := os.Getenv("DSN")

	e := echo.New()
	e.Use(middleware.CORS())

	sentry.Init(dsn)
	e.Logger.Fatal(e.Start(":" + port))
}
