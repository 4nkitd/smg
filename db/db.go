package db

import (
	"fmt"
	"strconv"
	"time"

	"smg/utils"

	"github.com/spf13/cobra"
)

var (
	currentTime = time.Now()
)

func DbDumper(cmd *cobra.Command, args []string) {

	fmt.Println("Creating a backup of all databases.")

	sst := strconv.Itoa(int(time.Now().Unix()))

	// dirname, _ := os.UserHomeDir()

	exportFile := sst + ".sql"

	utils.Exec("sudo", "mysqldump", "-u", "root", "--all-databases", ">", exportFile)

	fmt.Println("DataBase backup save to : " + exportFile)

}
