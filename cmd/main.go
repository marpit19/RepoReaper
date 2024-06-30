package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v62/github"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

type Config struct {
	GitHubToken    string `json:"github_token"`
	GitHubUsername string `json:"github_username"`
}

func loadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(byteValue, &config)
	return config, err
}

func main() {
	app := &cli.App{
		Name:  "Repo Reaper",
		Usage: "Delete multiple GitHub repositories",
		Action: func(c *cli.Context) error {
			config, err := loadConfig("config/config.json")
			if err != nil {
				log.Fatalf("Error loading config: %v", err)
			}

			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: config.GitHubToken},
			)
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)

			fmt.Print("Enter repository names to delete (comma-separated): ")
			var input string
			fmt.Scanln(&input)
			repositoriesToDelete := strings.Split(input, ",")

			for _, repo := range repositoriesToDelete {
				repo = strings.TrimSpace(repo)
				if repo == "" {
					continue
				}
				_, err := client.Repositories.Delete(ctx, config.GitHubUsername, repo)
				if err != nil {
					fmt.Printf("❌ Error deleting repository %s: %v\n", repo, err)
				} else {
					fmt.Printf("✅ Successfully deleted %s\n", repo)
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
