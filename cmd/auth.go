package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// AuthResponse represents the structure of the authentication response
type AuthResponse struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

var createCmd = &cobra.Command{
	Use:   "login",
	Short: "Authentication",
	Long:  "Authenticate with your credentials in the .env folder",

	// Run: func(cmd *cobra.Command, server, username, password string) (*AuthResponse, error) {

	// },
}

func loadCredentials() {
	godotenv.Load()
	authURL := os.Getenv("BaseURL") + "/api/token/"
	println(authURL)
}

// RetrieveTokens authenticates with the server and retrieves access and refresh tokens
func RetrieveTokens(server, username, password string) (*AuthResponse, error) {
	authURL := server + "/api/token/"

	// Prepare the JSON payload
	payload := map[string]string{
		"username": username,
		"password": password,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("failed to encode authentication payload: " + err.Error())
	}

	// Create the HTTP POST request with JSON payload
	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.New("failed to create authentication request: " + err.Error())
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("failed to send authentication request: " + err.Error())
	}
	defer resp.Body.Close()

	// Check if the response status code indicates success
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("authentication failed with status: " + resp.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response body: " + err.Error())
	}

	// Unmarshal the response JSON into the AuthResponse struct
	var authResponse AuthResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return nil, errors.New("failed to parse authentication response: " + err.Error())
	}

	return &authResponse, nil
}

// godotenv.Load()
// 	fmt.Println(os.Getenv("Password"))
