package internal

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/getToken", GetB2BTokenController)
	e.GET("/generateSignature", GenerateSignatureServiceController)
}
