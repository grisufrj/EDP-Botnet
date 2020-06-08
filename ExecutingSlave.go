package main


import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Request struct {
    Command string      `json:"command"`
    Args string            `json:"args"`
}

/*
func exec_cmd(string[] args) (byte[],int){

	fmt.Println(args)
	return [1,2],0


}
*/

func recv_cmd(url string)(Request){
    data := Request{}
    resp, atTheDisco := http.Get(url)
    if atTheDisco!= nil{
      panic(atTheDisco)
    }
    body, _ := ioutil.ReadAll(resp.Body)
    bodys:=string(body)

    json.Unmarshal([]byte(bodys), &data)
    return data
}


func main() {
	data:=recv_cmd("http://127.0.0.1:5000/command")
    	fmt.Println(data)
	
}
