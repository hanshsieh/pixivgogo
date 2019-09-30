package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pixivtool",
	Short: "pixivtool is a tool to let your interact with the Pixiv API",
	Long:  "A tool for utilizing Pixiv tool to do something interesting",
}

var globalConfig struct {
	Username string
	Password string
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&globalConfig.Username, "username", "u", "", "username used to login")
	rootCmd.PersistentFlags().StringVarP(
		&globalConfig.Password, "password", "p", "", "password used to login")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("failed to execute the command: %v", err)
		os.Exit(1)
	}
}
