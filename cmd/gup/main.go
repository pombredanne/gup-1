package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/r-medina/gup"
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
		fmt.Printf("updating %s...\n", execName)
		if err := run(execName); err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
		}
	}

}

func run(bin string) error {
	pkg, err := gup.GetPkg(bin)
	if err != nil {
		return err
	}
	fmt.Println(pkg)
	cmd := exec.Command("go", "get", "-u", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("go", "install", pkg)

	return cmd.Run()
}
