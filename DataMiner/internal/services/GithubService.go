package services

import (
	"DataMiner/internal/repository"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GithubService struct {
	GithubRepository repository.GithubRepository
}

func NewGithubService(githubRepository repository.GithubRepository) *GithubService {
	return &GithubService{GithubRepository: githubRepository}
}

func (s *GithubService) GetDataFromAPI(token string, username string) {
	query := map[string]string{"query": fmt.Sprintf(`query {
            user(login: "%s") {
              name
              contributionsCollection(from: "%s", to: "%s") {
                contributionCalendar {
                  totalContributions
                  weeks {
                    contributionDays {
                      contributionCount
                      date
                      weekday
                    }
                  }
                }
              }
            }
          }`, username, "2023-04-01T12:19:51Z", time.Now().Format(time.RFC3339)),
	}
	b, err := json.Marshal(&query)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Couldn't have made a request to github's graphql api : %v", err)
	}
	req.Header.Add("Authorization", "bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Failed to make a call to the API: %v", err)
	}
	scanner := bufio.NewScanner(resp.Body)
	body := ""
	for scanner.Scan() {
		body += scanner.Text()
	}
	fmt.Println(body)
}
