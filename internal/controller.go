package internal

import (
	"time"

	"github.com/labstack/echo/v4"
)

func GetB2BTokenController(c echo.Context) error {
	url := c.QueryParam("url")
	headerTimestamp := c.Request().Header.Get("X-TIMESTAMP")
	timestamp, err := time.Parse(SnapTimeFormat, headerTimestamp)
	if err != nil {
		return c.JSON(400, "Wrong timestamp")
	}

	request := GetB2BTokenRequest{
		Url:       url,
		Timestamp: timestamp,
	}

	response, err := GetB2BToken(&request)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, response)
}
