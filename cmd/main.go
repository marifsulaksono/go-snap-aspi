package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-snap-aspi/internal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("failed to load file .env")
	}

	e := echo.New()
	internal.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
