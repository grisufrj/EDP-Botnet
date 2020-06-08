package main

import (
	"encoding/json"
	f "fmt"
	"io/ioutil"
	"net/http"
)

type responseCommand struct {
	command string `json:"command"`
	path    string `json:"path"`
}

func main() {
	RespJson := responseCommand{}
	resp, _ := http.Get("http://127.0.0.1:5000/command")
	body, _ := ioutil.ReadAll(resp.Body)
	//	body := string(body)
	f.Println([]byte(string(body)))
	if err := json.Unmarshal(body, &RespJson); err != nil {
		f.Println(err)
	}
	f.Println(RespJson)

}
