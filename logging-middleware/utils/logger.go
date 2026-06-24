package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var BearerToken string

type logPayload struct {
	Stack   string `json:"stack"`
	Level   string `json:"level"`
	Package string `json:"package"`
	Message string `json:"message"`
}

func AuthenticateClient() error {
	baseURL := os.Getenv("AFFORDMED_BASE_URL")

	requestBody, _ := json.Marshal(map[string]string{
		"email":        os.Getenv("EMAIL"),
		"name":         os.Getenv("NAME"),
		"rollNo":       os.Getenv("ROLL_NO"),
		"accessCode":   os.Getenv("ACCESS_CODE"),
		"clientID":     os.Getenv("AFFORDMED_CLIENT_ID"),
		"clientSecret": os.Getenv("AFFORDMED_CLIENT_SECRET"),
	})

	response, err := http.Post(baseURL+"/evaluation-service/auth", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return err
	}

	if token, exists := result["access_token"]; exists {
		BearerToken = token.(string)
	}

	return nil
}

func Log(stack string, level string, packageName string, message string) {
	allowedPackages := map[string]bool{
		"cache": true, "controller": true, "cron_job": true, "db": true,
		"domain": true, "handler": true, "repository": true, "route": true,
		"service": true, "auth": true, "config": true, "middleware": true, "utils": true,
	}
	if !allowedPackages[packageName] {
		log.Printf("Local Error: Attempted to log with an invalid package name: %s", packageName)
		return
	}

	if BearerToken == "" {
		log.Printf("Local Error: No BearerToken available. Call AuthenticateClient() first.")
		return
	}

	payload := logPayload{
		Stack:   stack,
		Level:   level,
		Package: packageName,
		Message: message,
	}

	body, _ := json.Marshal(payload)
	baseURL := os.Getenv("AFFORDMED_BASE_URL")
	if baseURL == "" {
		baseURL = "http://4.224.186.213"
	}

	request, _ := http.NewRequest("POST", baseURL+"/evaluation-service/logs", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+BearerToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Log POST failed: %v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		fmt.Println("\n[✅ LOCAL SUCCESS] Log successfully pushed to Affordmed Database!")
	} else {
		fmt.Printf("\n[❌ LOCAL ERROR] Affordmed rejected the log. Status Code: %d\n", response.StatusCode)
	}
}
