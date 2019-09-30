package cmd

import (
	"github.com/spf13/cobra"
)

var illustCmd = &cobra.Command{
	Use:   "illust",
	Short: `"illust" allows you to do download related to illustrations`,
	Long:  `"illust" allows you to do operations related to illustrations`,
}

func init() {
	rootCmd.AddCommand(illustCmd)
}
