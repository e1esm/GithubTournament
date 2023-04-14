package main

import (
	"DataMiner/internal/repository"
	"DataMiner/internal/services"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	token := os.Getenv("GITHUB_TOKEN")
	githubRepo := repository.NewGithubRepository(nil)
	githubService := services.NewGithubService(*githubRepo)
	githubService.GetDataFromAPI(token, "e1esm")

}
