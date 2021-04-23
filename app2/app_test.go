package main

import (
	"bytes"
	"strings"
	"testing"

	cmder "github.com/yaegashi/cobra-cmder"
)

func TestApp(t *testing.T) {
	testCases := []struct {
		args []string
		want string
		err  bool
	}{
		{args: []string{"calc", "add", "-a", "1", "-b", "2"}, want: "3\n", err: false},
		{args: []string{"calc", "sub", "-a", "1", "-b", "2"}, want: "-1\n", err: false},
		{args: []string{"calc", "add"}, want: "Error: missing flags -a and/or -b\n", err: true},
		{args: []string{"calc", "sub"}, want: "Error: missing flags -a and/or -b\n", err: true},
	}
	for _, tC := range testCases {
		args := strings.Join(tC.args, " ")
		t.Run(args, func(t *testing.T) {
			buf := &bytes.Buffer{}
			cmd := cmder.Cmd(&App{Out: buf})
			cmd.SetArgs(tC.args)
			cmd.SetOut(buf)
			cmd.SetErr(buf)
			err := cmd.Execute()
			if tC.err && err == nil {
				t.Errorf("%q returns no error", args)
			}
			if !tC.err && err != nil {
				t.Errorf("%q returns error: %s", args, err)
			}
			got := buf.String()
			if got != tC.want {
				t.Errorf("%q returns %q, want %q", args, got, tC.want)
			}
		})
	}
}
