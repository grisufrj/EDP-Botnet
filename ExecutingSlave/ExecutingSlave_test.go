package slave_test

import (
	"log"
	"os"
	"os/exec"
	"testing"

	slave "github.com/grisufrj/EDP-Botnet/ExecutingSlave"
)

func TestExecCmd(t *testing.T) {
	l := log.New(os.Stdout, "slave-logger", log.LstdFlags)
	s := slave.NewSlave(l)

	want := "test"
	cmd := exec.Command("echo", "-n", "test")

	outBuffer := s.ExecCmd(cmd)
	got := outBuffer.String()

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}
