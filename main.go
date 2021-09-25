package main

import (
	"github.com/spf13/cobra"

	cmd "smg/cmd"
)

var (
	AppName = "smg"
	RootCmd = &cobra.Command{Use: AppName}
)

func main() {

	RootCmd.AddCommand(cmd.InitServer())
	RootCmd.AddCommand(cmd.DbDumper())

	RootCmd.Execute()

}
