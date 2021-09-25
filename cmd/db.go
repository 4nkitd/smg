package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"smg/db"
)

func DbDumper() *cobra.Command {

	c := &cobra.Command{
		Use:   "dbDump filename.sql",
		Short: "dbDump filename.sql",
		Long:  `"dbDump" : This Command allow you to take backup to all databases.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   db.DbDumper,
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Annotations["error"])
		},
	}

	return c

}
