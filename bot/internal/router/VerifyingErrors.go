package router

import (
	"fmt"
	"github.com/enescakir/emoji"
)

type VerifyingErrors int

const (
	RequestError VerifyingErrors = iota
	InvalidUsername
	AlreadyHaveAccount
	BusyAccount
)

var errorsMap map[VerifyingErrors]string

func init() {
	errorsMap = map[VerifyingErrors]string{
		RequestError:       "Couldn't make a request to github, try once again later.",
		InvalidUsername:    "Account with this username wasn't found",
		AlreadyHaveAccount: fmt.Sprintf("You already have linked account %v\n", emoji.RedCircle),
		BusyAccount:        fmt.Sprintf("This account's already linked %v\n", emoji.RedCircle),
	}
}
