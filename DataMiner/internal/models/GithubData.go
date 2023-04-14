package models

import "DataMiner/internal/utils"

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var contributions []*Contribution
var currentIndex = 0

func init() {
	contributions = make([]*Contribution, 0, 100)
}

type GithubAccount struct {
	Username      string
	Contributions []Contribution
}

func NewFetchedGithubAccount(Username string, contributions []Contribution) *GithubAccount {
	return &GithubAccount{Username: Username, Contributions: contributions}
}

type Contribution struct {
	Amount  int
	Date    string
	Weekday string
}

func NewContribution(Amount int, Date string, weekday int) {
	contributions[currentIndex] = &Contribution{Amount: Amount, Date: utils.GetDayInString(Weekday(weekday))}
	currentIndex++
}
