package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func JsonHandleFunc(rw http.ResponseWriter, req *http.Request) {
	rw = GetHead(rw)
	if req.Method == "OPTIONS" {
		s := req.Header.Get("Access-Control-Request-Headers")
		if strings.Contains(s, "authorization") == true || strings.Contains(s, "Authorization") == true {
			rw.WriteHeader(204)
		}
		rw.WriteHeader(400)
		return
	}
	var output Replay
	if req.Method == "GET" {
		data := FindTask("./recode.json")
		json.NewEncoder(rw).Encode(data)
		return
	}
	body, e := ioutil.ReadAll(req.Body)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	if req.Method == "POST" {
		input := Replay{}
		e = json.Unmarshal(body, &input)
		output = input
		Recodes(input)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		json.NewEncoder(rw).Encode(output)
		fmt.Printf("%#v\n", input)
		return
	}
}
