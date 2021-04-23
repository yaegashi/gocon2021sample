package main

import (
	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

// AppCalcAdd - app calc add コマンド
type AppCalcAdd struct {
	*AppCalc // 親コマンド埋め込み
}

// NewAppCalcAdd - AppCalcAdd 作成
func (app *AppCalc) NewAppCalcAdd() cmder.Cmder { return &AppCalcAdd{AppCalc: app} }

// Cmd - cmder.Cmder インターフェース実装
func (app *AppCalcAdd) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Calculate A + B",
		RunE:  app.RunE,
	}
	return cmd
}

// RunE - メインルーチン
func (app *AppCalcAdd) RunE(cmd *cobra.Command, args []string) error {
	app.Println(app.A + app.B)
	return nil
}
