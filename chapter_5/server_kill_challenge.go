/*
	-- Server Kill Challenge --

- Write a function called killServer(pidFile string) error function that reads a process identifer from pidFile,
	converts it to an integer, and prints "killing <pid>" (instead of using os.Kill)
- Use github.com/pkg/errors to wrap errors
- Use io/ioutil ReadFile to read the file
- Use strconv.Atoi to convert the file content to an integer
*/
package main

import (
	"io/ioutil"
	"strconv"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	// get contents of the file
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "could not read file")
	}

	if err := os.Remove(pidFile); err != nil {
		// we can go on if we fail here
		log.Printf("warning: can't remove pid file %s", err)
	}

	// convert the pid number to an integer
	strPid := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPid))
	if err != nil {
		return errors.Wrap(err, "could not convert PID to an integer")
	}

	// Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprint(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
