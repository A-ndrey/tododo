package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/A-ndrey/tododo/internal/config"
	"log"
	"net/http"
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
	protocol := "http"
	if cfg.Port == 443 {
		protocol = "https"
	}
	url := fmt.Sprintf("%s://%s:%d%s/user?service=%s", protocol, cfg.Host, cfg.Port, apiPath, cfg.Service)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", token)

	client := http.Client{}

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
