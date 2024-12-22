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
	rootCmd        *cobra.Command
	CurrentCommand string
	HandleInit     tui.HandleInitFunc
}

var rootHolder = &CommandHolder{
	rootCmd: &cobra.Command{
		Use:   "start",
		Short: "A time tracker TUI",
		Long:  "A basic time tracker application for personal use",
	},
}

func Execute() *CommandHolder {
	rootHolder.rootCmd.SetUsageFunc(boa.UsageFunc)
	rootHolder.rootCmd.SetHelpFunc(boa.HelpFunc)
	err := rootHolder.rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return rootHolder
}

func init() {
	rootHolder.rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
