/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	boa "github.com/elewis787/boa"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "A time tracker TUI",
	Long:  "A basic time tracker application for personal use",
}

func Execute() *cobra.Command {
	rootCmd.SetUsageFunc(boa.UsageFunc)
	rootCmd.SetHelpFunc(boa.HelpFunc)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return rootCmd
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
