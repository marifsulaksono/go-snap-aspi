package internal

import (
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
)

func GetB2BTokenController(c echo.Context) error {
	headerURL := c.Request().Header.Get("X-PARTNER-URL")
	headerTimestamp := c.Request().Header.Get("X-TIMESTAMP")
	timestamp, err := time.Parse(SnapTimeFormat, headerTimestamp)
	if err != nil {
		return c.JSON(400, "Wrong timestamp")
	}

	request := GetB2BTokenRequest{
		Url:       headerURL,
		Timestamp: timestamp,
	}

	response, err := GetB2BToken(&request)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, response)
}

func GenerateSignatureServiceController(c echo.Context) error {
	headerTimestamp := c.Request().Header.Get("X-TIMESTAMP")
	endpointURL := c.Request().Header.Get("X-ENDPOINT-URL")
	endpointMethod := c.Request().Header.Get("X-ENDPOINT-METHOD")
	authorization := c.Request().Header.Get("Authorization")

	timestamp, err := time.Parse(SnapTimeFormat, headerTimestamp)
	if err != nil {
		return c.JSON(400, "Wrong timestamp")
	}

	var payload any
	err = c.Bind(&payload)
	if err != nil {
		return c.String(400, "bad request")
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return c.JSON(500, "failed")
	}

	request := GenerateSignatureServiceRequest{
		Timestamp:      timestamp,
		EndpointURL:    endpointURL,
		EndpointMethod: endpointMethod,
		Token:          authorization,
		Body:           body,
	}

	response, err := GenerateSignatureService(&request)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, response)
}
