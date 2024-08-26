/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/vladl2c/pixel_trigger/pkg/trigger"
)

const logo = `
 ____   _             _        ____                      _   
|  _ \ (_)__  __ ___ | |      / ___|   ___  ___   _   _ | |_ 
| |_) || |\ \/ // _ \| | _____\___ \  / __|/ _ \ | | | || __|
|  __/ | | >  <|  __/| ||_____|___) || (__| (_) || |_| || |_ 
|_|    |_|/_/\_\\___||_|      |____/  \___|\___/  \__,_| \__|
`

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

const (
	startAction    = "Start"
	settingsAction = "Settings"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", logoStyle.Render(logo))
		fmt.Printf("%s\n", logoStyle.Render("Running..."))
		fmt.Printf("%s\n", logoStyle.Render("Press ctrl + c to exit"))
		bot := trigger.Init()
		bot.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
