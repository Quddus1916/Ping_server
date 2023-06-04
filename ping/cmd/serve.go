/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	serveCmd.PersistentFlags().IntP("p", "p", 8080, "port where it will run")
	viper.BindPFlag("app.http_port", serveCmd.PersistentFlags().Lookup("p"))
	rootCmd.AddCommand(serveCmd)
}
