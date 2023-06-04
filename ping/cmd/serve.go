/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"

	"fmt"
	"github.com/ping/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
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
	Run: serve,
}

func init() {
	serveCmd.PersistentFlags().IntP("p", "p", 8080, "port where it will run")
	viper.BindPFlag("app.http_port", serveCmd.PersistentFlags().Lookup("p"))
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	appCfg := config.AppCfg()
	fmt.Println(appCfg)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	htsrvr := &http.Server{
		ReadTimeout:  appCfg.ReadTimeout,
		WriteTimeout: appCfg.WriteTimeout,
		IdleTimeout:  appCfg.IdleTimeout,
		Addr:         ":" + strconv.Itoa(appCfg.HTTPPort),
		//Handler:      router.Route(),
	}

	go func() {
		log.Println("HTTP: Listening on port " + strconv.Itoa(appCfg.HTTPPort))
		log.Fatal(htsrvr.ListenAndServe())
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	log.Println("Shutting down server...")
	htsrvr.Shutdown(ctx)
	log.Println("Server shutdown gracefully")
	cancel()
}
