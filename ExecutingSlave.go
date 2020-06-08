package main

import (
	//	"encoding/json"
	f "fmt"
	"net/http"
	//"io/ioutil"
)

type responseCommand struct {
	command string
	path    string
}

func main() {
	//	var RespJson responseCommand
	resp, _ := http.Get("http://127.0.0.1:5000/command")
	f.Println(resp)
	f.Println(resp.Body)
	//	err := json.NewDecoder(resp.Body).Decode(&RespJson)

	//if err != nil {
	//		f.Println(err)
	//	}
	//	f.Println(RespJson)

}
