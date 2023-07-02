package commons

import "log"

func AppRecover() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}
