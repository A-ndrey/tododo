package services

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/A-ndrey/tododo/internal/config"
	"log"
	"net"
	"net/http"
	"time"
)

const apiPath = "/api/v1"

type AuthService interface {
	GetUserEmail(token string) (string, error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) GetUserEmail(token string) (string, error) {
	cfg := config.GetAuth()

	url := fmt.Sprintf("%s%s/user?service=%s", cfg.Host, apiPath, cfg.Service)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", token)

	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("auth-service response status not OK")
	}

	jsonEmail := struct {
		Email string `json:"email"`
	}{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&jsonEmail)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return jsonEmail.Email, nil
}
