package main

import (
	"os"

	cmder "github.com/yaegashi/cobra-cmder"
)

func main() {
	app := &App{Out: os.Stdout}
	cmd := cmder.Cmd(app)
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
