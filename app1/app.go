package main

import (
	"github.com/spf13/cobra"
)

// App - app コマンド (トップレベル)
type App struct {
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
