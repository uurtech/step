package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	path := usr.HomeDir + "/.step/"
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

var listCmd = &cobra.Command{
	Use:              "list",
	Short:            "List your favorite SSH commands",
	Long:             `List your favorite SSH commands`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			panic(err)
		}
		rows, err := db.Query("SELECT * FROM commands")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var alias string
			var command string
			err = rows.Scan(&id, &alias, &command)
			if err != nil {
				panic(err)
			}
			fmt.Println("==============================================================================")
			fmt.Printf("%d \t alias: %s - Command: %s\n", id, alias, command)
			fmt.Println("==============================================================================")
		}
	},
}

var initCMD = &cobra.Command{
	Use:   "init",
	Short: "Initialize your favorite SSH commands",
	Long:  `Initialize your favorite SSH commands`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		path := usr.HomeDir
		shCmd := exec.Command("sh", "-c", "echo \"$SHELL\"")
		output, _ := shCmd.CombinedOutput()
		shell := string(output)
		tempPath := path + "/.step/steps.sh"

		if _, err := os.Stat(tempPath); os.IsNotExist(err) {
			f, err := os.Create(tempPath)
			if err != nil {
				panic(err)
			}
			defer f.Close()
		}

		switch shell {
		case "/bin/zsh\n":

			f, err := os.OpenFile(path+"/.zshrc", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				panic(err)
			}

			defer f.Close()
			if _, err := f.Write([]byte("source " + tempPath)); err != nil {
				log.Fatal(err)
			}
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		}

	},
}

var loadCMD = &cobra.Command{
	Use:              "load",
	Short:            "Load your favorite SSH commands",
	Long:             `Load your favorite SSH commands`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			panic(err)
		}
		rows, err := db.Query("SELECT * FROM commands")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var alias string
			var command string
			err = rows.Scan(&id, &alias, &command)
			if err != nil {
				panic(err)
			}
			fmt.Println("==============================================================================")
			fmt.Printf("%d \t alias: %s - Command: %s\n", id, alias, command)
			fmt.Println("==============================================================================")
		}

		//create file

	},
}
