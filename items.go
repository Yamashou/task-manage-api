package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
  "os"
  "log"
)

func GetHead(rw http.ResponseWriter)http.ResponseWriter{
  rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Credentials", "true")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	rw.Header().Set("Content-Type", "application/json")
  return rw
}

func Recodes(recorde Replay) {
	fout, err := os.Create("recode.json")
	if err != nil {
		fmt.Println("Recode: ", err)
	}
	outputJson, err := json.Marshal(&recorde)
	fout.Write([]byte(outputJson))
	if err != nil {
		panic(err)
	}
	defer fout.Close()
	fmt.Println("Recoding file")
}

func FindTask(task string) Replay {
	bytes, err := ioutil.ReadFile(task)
	if err != nil {
		log.Fatal(err)
	}
	var data Replay
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	return data
}
