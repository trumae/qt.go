package main

import (
	"os/exec"

	"github.com/kitech/qt.go/qtrt"
)

func runcmdout(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	occ, err := cmd.Output() // Output runs the command and returns its standard output.
	qtrt.ErrPrint(err, name, arg)
	if err != nil {
		return "", err
	}
	cmd.Wait()
	return string(occ), nil
}
