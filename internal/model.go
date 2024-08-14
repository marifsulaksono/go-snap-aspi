package internal

import "time"

const SnapTimeFormat string = "2006-01-02T15:04:05-07:00"

type GetB2BTokenRequest struct {
	Url       string
	Timestamp time.Time
}

type GetB2BTokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   string `json:"expiresIn"`
	TokenType   string `json:"tokenType"`
}
