package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bfontaine/which/which"
	"github.com/pkg/errors"
)

const usage = `gup [binary ...]

gup updates go binaries. It finds the go import path of the binary and does 'go get -u' and 'go install' for that path if found.
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		return
	}

	for _, execName := range os.Args[1:] {
		fmt.Printf("updating %q...\n", execName)
		if err := run(execName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

func run(execName string) error {
	path, err := getExecPath(execName)
	if err != nil {
		return err
	}

	src, err := getMainPath(path)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "get", "-u", src)
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("go", "install", src)

	return cmd.Run()
}

//
// everything below this point is from github.com/christophberger/goman
//

// Copyright (c) 2017, Christoph Berger
// All rights reserved.

// getExecPath receives the name of an executable and determines its path
// based on $PATH, $GOPATH, or the current directory (in this order).
func getExecPath(name string) (string, error) {

	// Try $PATH first.
	s := which.One(name) // $ which <name>
	if s != "" {
		return s, nil
	}

	// Next, try $GOPATH/bin
	paths := gopath()
	for i := 0; s == "" && i < len(paths); i++ {
		s = which.OneWithPath(name, paths[i]+filepath.Join("bin"))
	}
	if s != "" {
		return s, nil
	}

	// Finally, try the current directory.
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "Unable to determine current directory")
	}
	s = which.OneWithPath(name, wd)
	if s == "" {
		return "", errors.New(name + " not found in any of " + os.Getenv("PATH") + ":" + strings.Join(paths, ":"))
	}

	return s, nil
}

// gopath returns a list of paths as defined by the GOPATH environment
// variable, or the default gopath if $GOPATH is empty.
func gopath() []string {

	gp := os.Getenv("GOPATH")
	if gp == "" {
		return []string{build.Default.GOPATH}
	}

	return strings.Split(gp, pathssep())
}

// pathssep returns the separator between the paths of $PATH or %PATH%.
func pathssep() string {

	sep := ":"
	if runtime.GOOS == "windows" {
		sep = ";"
	}

	return sep
}
