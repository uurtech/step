/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "step",
	Short: "Creates alias for your favorite SSH commands",
	Long: `You can create alias for your favorite SSH and SCP commands:
		No need to Search in Bash History !
		`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		-h			help
		store			Saves the alias for the spesified ssh
		update			Save alias with a key path
		command 		define the command to be executed with the alias	
		============
		Create Alias
		============
		step store demo_1 -c "ssh user@host"
	`)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.step.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(storeCmd)
	rootCmd.AddCommand(listCmd)
}
