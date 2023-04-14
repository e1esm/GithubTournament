package main

import (
	"DataMiner/internal/repository"
	"DataMiner/internal/services"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
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

	s := gocron.NewScheduler(time.UTC)
	s.Every("1m").Do(func() {
		fmt.Println("hello world")
	})
	s.StartBlocking()
}
