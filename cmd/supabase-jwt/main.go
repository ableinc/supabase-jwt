package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/supabase-community/supabase-go"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v3"
)

// Define the structure that matches the configuration YAML file
type Config struct {
	AppName           string `yaml:"APP_NAME"`
	SupbaseProjectUrl string `yaml:"SUPABASE_PROJECT_URL"`
	SupabaseApiKey    string `yaml:"SUPABASE_API_KEY"`
}

func getUserCredentials() (string, string) {
	var email string
	var password string
	// Get User's email
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter your password: ")
	// Get User's password - Turn off input echoing
	bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("Error reading password: %v", err)
	}
	password = string(bytePassword)
	return email, password
}

func supabaseLogin(config Config, emailAddress string, password string) {
	fmt.Printf("\nInitializing Supabase connection for %s project...\n", config.AppName)
	client, err := supabase.NewClient(config.SupbaseProjectUrl, config.SupabaseApiKey, nil)
	if err != nil {
		log.Fatalf("Unable to initialize connection to Supabase. Reason: %v", err)
	}
	token, err := client.Auth.SignInWithEmailPassword(emailAddress, password)
	if err != nil {
		log.Fatalf("Unable to login with credentials provided. Reason: %v", err)
	}
	fmt.Printf("Hi %s, Successfully logged in! Your token is below:\n\n", token.User.Email)
	fmt.Println(token.AccessToken)
}

func main() {
	// Read YAML file
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to get current working directory: %v", err)
	}
	var configPath string = filepath.Join(cwd, "configs", "config.yml")
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}
	// Unmarshall the YAML file into the Config struct
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshall configuration file: %v", err)
	}
	if config.AppName == "" || config.SupabaseApiKey == "" || config.SupbaseProjectUrl == "" {
		log.Fatal("You must update configs/config.yml file and provide all fields.")
	}
	// User message
	fmt.Printf("[!] ATTENTION: You are creating an access token for the %s project [!]\n\n", config.AppName)
	// Prompt user for credentials
	email, password := getUserCredentials()
	// Login to supabase and get access token
	supabaseLogin(config, email, password)
}
