package main

import (
	"fmt"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

// AppCalc - app calc コマンド
type AppCalc struct {
	*App     // 親コマンド埋め込み
	A    int // -a フラグの値
	B    int // -b フラグの値
}

// NewAppCalc - AppCalc 作成
func (app *App) NewAppCalc() cmder.Cmder { return &AppCalc{App: app} }

// Cmd - cmder.Cmder インターフェース実装
func (app *AppCalc) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "calc",
		Short:             "Calculate A and B",
		PersistentPreRunE: app.PersistentPreRunE,
	}
	cmd.PersistentFlags().IntVarP(&app.A, "value-a", "a", 0, "value A")
	cmd.PersistentFlags().IntVarP(&app.B, "value-b", "b", 0, "value B")
	return cmd
}

// PersistentPreRunE - フラグ存在チェック
func (app *AppCalc) PersistentPreRunE(cmd *cobra.Command, args []string) error {
	if !cmd.Flags().Changed("value-a") || !cmd.Flags().Changed("value-b") {
		return fmt.Errorf("missing flags -a and/or -b")
	}
	return nil
}
