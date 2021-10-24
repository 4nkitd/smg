package main

import (
	"fmt"
	"os"
	"smg/cmd"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	AppName = "smg"
	RootCmd = &cobra.Command{Use: AppName}
	cfgFile string
)

type Header struct {
	key   string
	value string
}

func main() {

	RootCmd.AddCommand(cmd.InitServer())
	RootCmd.AddCommand(cmd.DbDumper())
	RootCmd.AddCommand(cmd.Decrypt())
	RootCmd.AddCommand(cmd.Encrypt())
	RootCmd.AddCommand(cmd.Gen())
	RootCmd.AddCommand(cmd.GenKey())
	RootCmd.AddCommand(cmd.Gen())
	RootCmd.AddCommand(cmd.Verify())

	RootCmd.Execute()

}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".EDH")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
