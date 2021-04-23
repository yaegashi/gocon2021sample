package main

import (
	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

// AppCalcSub - app calc sub コマンド
type AppCalcSub struct {
	*AppCalc // 親コマンド埋め込み
}

// NewAppCalcSub - AppCalcSub 作成
func (app *AppCalc) NewAppCalcSub() cmder.Cmder { return &AppCalcSub{AppCalc: app} }

// Cmd - cmder.Cmder インターフェース実装
func (app *AppCalcSub) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sub",
		Short: "Calculate A - B",
		RunE:  app.RunE,
	}
	return cmd
}

// RunE - メインルーチン
func (app *AppCalcSub) RunE(cmd *cobra.Command, args []string) error {
	app.Println(app.A - app.B)
	return nil
}
