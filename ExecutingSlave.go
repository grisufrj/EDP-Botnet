package main

import (
	"encoding/json"
	f "fmt"
	"net/http"
	"io/ioutil"
)

type responseCommand struct {
	command string
	path    string
}

func main() {
	RespJson :=responseCommand{}
	resp, _ := http.Get("http://127.0.0.1:5000/command")
	body, _ := ioutil.ReadAll(resp.Body)
	f.Println(string(body))
	json.Unmarshal([]byte(string(body)),&RespJson)
	f.Println(RespJson.command)

	err:=json.NewDecoder(resp.Body).Decode(&RespJson)

			f.Println(err)
	f.Println(RespJson)

}
