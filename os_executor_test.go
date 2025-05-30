package command

import (
	"io"
	"runtime"
	"testing"

	"github.com/i9si-sistemas/assert"
	"github.com/i9si-sistemas/command/spy"
)

func TestCommandExecutor(t *testing.T) {
	stdout := spy.New()
	stderr := spy.New()

	var (
		runner   Runner
		expected string
	)
	if runtime.GOOS == "windows" {
		runner, expected = testWindows(stdout, stderr)
	} else {
		runner, expected = testUnix(stdout, stderr)
	}

	err := runner.Run()
	assert.Nil(t, err)
	assert.Empty(t, stderr.Data())

	assert.Equal(t, string(stdout.Data()), expected)
}

func testWindows(stdout, stderr io.Writer) (runner Runner, expectedOutput string) {
	echoPath := "C:\\Windows\\System32\\cmd.exe"
	runner = New(NewRunner(stdout, stderr)).WithPath(echoPath).AppendArgs("/C", "echo", "hello world")
	expectedOutput = "\"hello world\"\r\n"
	return
}

func testUnix(stdout, stderr io.Writer) (runner Runner, expectedOutput string) {
	echoPath := "/bin/echo"
	runner = New(NewRunner(stdout, stderr)).WithPath(echoPath).AppendArgs("echo", "hello world")
	expectedOutput = "hello world\n"
	return
}
