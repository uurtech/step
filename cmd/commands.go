package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/user"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var dbPath string
var Command string

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	path := usr.HomeDir + "/step/"
	dbPath = path + "data.db"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	if _, err := os.Stat(dbPath); err != nil || os.IsNotExist(err) {
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			panic(err)
		}
		// we are creating sqlite database/table
		_, err = db.Exec("CREATE TABLE `commands` (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`alias` VARCHAR(255) NULL,`command` VARCHAR(255) NULL)")
		if err != nil {
			log.Fatal(err)
		}
		db.Close()
	}

	// have subcommand
	// storeCmd.AddCommand(commandCmd)
	storeCmd.Flags().StringVarP(&Command, "command", "c", "", "pass command to execute")

}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Step",
	Long:  `Love open source !`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Step v0.1.0")
	},
}

var storeCmd = &cobra.Command{
	Use:              "store",
	Short:            "Store your favorite SSH commands",
	Long:             `Store your favorite SSH commands`,
	TraverseChildren: true,
	Args:             cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if Command == "" {
			fmt.Println("Command is required")
			return
		}
		alias := args[0]

		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			panic(err)
		}
		stmt, err := db.Prepare("INSERT INTO commands(alias, command) values(?,?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(alias, Command)
		if err != nil {
			panic(err)
		}
		fmt.Println("Command stored")
	},
}
