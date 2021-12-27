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
	}
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
	Use:   "store",
	Short: "Store your favorite SSH commands",
	Long:  `Store your favorite SSH commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", args)
		_, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			panic(err)
		}
		// _, err := db.Prepare("INSERT INTO remotes(alias, keypath, machine) values(?,?,?)")
		// if err != nil {
		// 	panic(err)
		// }
		// _, err = stmt.Exec(remote.Alias, remote.KeyPath, remote.Machine)
		// if err != nil {
		// 	panic(err)
		// }
	},
}
