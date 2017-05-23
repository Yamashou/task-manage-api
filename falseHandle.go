package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func FalseHandler(rw http.ResponseWriter, req *http.Request) {
	rw = GetHead(rw)
	if req.Method == "OPTIONS" {
		s := req.Header.Get("Access-Control-Request-Headers")
		if strings.Contains(s, "authorization") == true || strings.Contains(s, "Authorization") == true {
			rw.WriteHeader(204)
		}
		rw.WriteHeader(400)
		return
	}
	if req.Method != "GET" {
		return
	}
	if req.Method == "GET" {
		data := FindTask("./recode.json")
		replay := Replay{}
		for _, v := range data.Recode {
			if !v.Done {
				replay.Recode = append(replay.Recode, v)
			}
		}
		json.NewEncoder(rw).Encode(replay)
		return
	}
}
