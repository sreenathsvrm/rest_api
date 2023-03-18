package db

import (
	"time"
)

type Userdetails struct {
	Name string
	Pass string
}

var Users = make(map[string]Userdetails)

// session
type Session struct {
	Name   string
	Expire time.Time
}

var SessionToken string

func (s Session) Sessionexpired() bool {

	return s.Expire.Before(time.Now())
}

var Sessions = make(map[string]Session)

type Messages struct {
	Color   string
	Message string
}

// this have the message what we want to show in login
var LoginMessage = Messages{}
