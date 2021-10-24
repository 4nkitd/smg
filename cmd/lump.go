package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	lump "smg/lump"
)

var (
	Pass string
)

func InitServer() *cobra.Command {

	c := &cobra.Command{
		Use:   "lump [Nginx, Mysql, Php8]",
		Short: "lump nginx mysql php",
		Long: `"lump" Commands allow you to install you need to setup a new Server.
				
				# Supported Applications List

					1. Nginx
					2. Mysql
					3. Php7.4
					4. Php7.3
					5. Php7.2
		
		`,
		Args: cobra.MinimumNArgs(1),
		Run:  lump.InstallerInit,
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Annotations["error"])
		},
	}

	c.Flags().StringVar(&Pass, "pass", "password", "password for mysql installation.")

	return c

}
