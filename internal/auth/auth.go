package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/A-ndrey/tododo/internal/config"
	"log"
	"net/http"
)

const apiPath = "/api/v1"

func GetUserEmail(token string) (string, error) {
	cfg := config.GetAuth()
	url := fmt.Sprintf("http://%s:%d%s/user?service=%s", cfg.Host, cfg.Port, apiPath, cfg.Service)

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
