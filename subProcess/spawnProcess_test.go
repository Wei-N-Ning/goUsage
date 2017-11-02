package subProcess

import (
	"testing"
	"os/exec"
	"io/ioutil"
)


func TestGetTextFromStdOut(t *testing.T) {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()  // wait for command
	if err != nil || len(dateOut) < 1 {
		t.Error(err)
	}
}


func TestGetTextFromStdOutExpectError(t *testing.T) {
	dateCmd := exec.Command("date11")
	_, err := dateCmd.Output()
	if err == nil {
		t.Error()
	}
}


func TestSendTextToStdInExpectReceived(t *testing.T) {
	grepCmd := exec.Command("grep", "foo")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("asdasd asd foo\nfoo"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	t.Log(string(grepBytes))
}


func TestCallLsUsingDelineatedArgArray(t *testing.T) {
	lsCmd := exec.Command("ls", "-l", "-A")
	lsOut, err := lsCmd.Output()
	if err != nil {
		t.Error()
	}
	t.Log(string(lsOut))
}


func TestStartCommandExpectNonBlocking(t *testing.T) {
	sleepCmd := exec.Command("sleep", "1")
	sleepCmd.Start()
	t.Logf("process is done: %+v", sleepCmd)
	sleepCmd.Wait()
	t.Logf("process is done: %+v", sleepCmd)
}



