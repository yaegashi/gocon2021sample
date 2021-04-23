package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// App - app コマンド (トップレベル)
type App struct {
	Out io.Writer // 標準出力
}

// Cmd - cmder.Cmder インターフェース実装
func (app *App) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "app",
		Short:        "cobra-comder application",
		SilenceUsage: true, // エラー発生時のヘルプ出力を抑制
	}
	return cmd
}

func (app *App) Print(args ...interface{}) (int, error) {
	return fmt.Fprint(app.Out, args...)
}

func (app *App) Println(args ...interface{}) (int, error) {
	return fmt.Fprintln(app.Out, args...)
}

func (app *App) Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(app.Out, format, args...)
}
