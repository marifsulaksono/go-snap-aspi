package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetB2BToken(request *GetB2BTokenRequest) (response *GetB2BTokenResponse, err error) {
	stringToSign := fmt.Sprintf("%s|%s", os.Getenv("CLIENT_KEY"), request.Timestamp)
	signature, err := GenerateAsymmetricSignature(stringToSign)
	if err != nil {
		err = errors.New("error generate asymmetric signature")
		return
	}

	log.Println(signature)
	bodyPayload, _ := json.Marshal(map[string]string{
		"grantType": "client_credentials",
	})

	req, err := http.NewRequest(http.MethodPost, request.Url, bytes.NewBuffer(bodyPayload))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-SIGNATURE", signature)
	req.Header.Set("X-TIMESTAMP", request.Timestamp.String())
	req.Header.Set("X-CLIENT-KEY", os.Getenv("CLIENT_KEY"))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Body)
		err = errors.New("failed to get token b2b")
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return
	}

	return
}
