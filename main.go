package main

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"time"
)

type Data struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	DeadLine  time.Time `json:"deadline"`
	DoneTime  time.Time `json:"donetime"`
}

type Recode struct {
	Done bool `json:"done"`
	Data Data `json:"data"`
}

type Replay struct {
	Recode []Recode `json:"recode"`
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}
	http.HandleFunc("/", JsonHandleFunc)
	http.HandleFunc("/not", FalseHandler)
	fcgi.Serve(l, nil)
}
