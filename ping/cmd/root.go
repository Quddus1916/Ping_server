/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/ping/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your application",
	Long:  `A brief description of your application`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	config.Init()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		return
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
}
