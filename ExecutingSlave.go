package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type Request struct {
	//	Command string   `json:"command"`
	Args []string `json:"args"`
}

/*
func exec_cmd(string[] args) (byte[],int){

	fmt.Println(args)
	return [1,2],0


}
*/

func recv_cmd(url string) (Request, exec.Cmd) {
	data := Request{}
	resp, atTheDisco := http.Get(url)
	if atTheDisco != nil {
		panic(atTheDisco)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	bodys := string(body)

	json.Unmarshal([]byte(bodys), &data)
	data2 := exec.Cmd{
		Args: data.Args,
	}
	return data, data2
}

func main() {
	data, data2 := recv_cmd("http://127.0.0.1:5000/command")
	fmt.Println(data)
	//cmd := exec.Command(data.Command, data.Args)
	var out bytes.Buffer
	data2.Stdout = &out
	if err := data2.Run(); err != nil {
		fmt.Println(err)
	}
}
