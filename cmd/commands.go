package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Step",
	Long:  `All software has versions. This is Step's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Step v0.1.0")
	},
}

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store your favorite SSH commands",
	Long:  `Store your favorite SSH commands`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
