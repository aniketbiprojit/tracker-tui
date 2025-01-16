/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"tracker/cmd/tui"

	boa "github.com/elewis787/boa"
	"github.com/spf13/cobra"
)

type CommandHolder struct {
	RootCmd        *cobra.Command
	CurrentCommand string
	HandleInit     tui.HandleInitFunc
}

var RootHolder = &CommandHolder{
	RootCmd: &cobra.Command{
		Use:   "tracker",
		Short: "A time tracker TUI",
		Long:  "A basic time tracker application for personal use",
	},
}

func Execute() *CommandHolder {
	RootHolder.RootCmd.SetUsageFunc(boa.UsageFunc)
	RootHolder.RootCmd.SetHelpFunc(boa.HelpFunc)
	err := RootHolder.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return RootHolder
}

func init() {
	RootHolder.RootCmd.Flags().BoolP("help", "h", false, "Help message for toggle")
}
