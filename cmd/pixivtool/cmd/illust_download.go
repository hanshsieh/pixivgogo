package cmd

import (
	"log"

	"github.com/sleepingpig/pixivgogo/cmd/pixivtool/tool"
	"github.com/spf13/cobra"
)

func init() {
	downloadIllustConfig := &tool.DownloadIllustConfig{}
	var downloadIllustCmd = &cobra.Command{
		Use:   "download",
		Short: `"download" allows you to download illustrations`,
		Long:  `"download" allows you to download illustrations`,
		Run: func(cmd *cobra.Command, args []string) {
			pixivTool := tool.NewIllustsDownloader()
			downloadIllustConfig.Username = globalConfig.Username
			downloadIllustConfig.Password = globalConfig.Password
			if err := pixivTool.DownloadIllustrations(downloadIllustConfig); err != nil {
				log.Fatal(err)
			}
		},
	}
	downloadIllustCmd.Flags().StringVar(
		&downloadIllustConfig.DstDirectory, "output", "",
		"directory to download the files")
	downloadIllustCmd.Flags().IntVar(
		&downloadIllustConfig.Count, "count", 0,
		"number of illustrations to download. (An illustration may contain multiple images)")
	if err := downloadIllustCmd.MarkFlagRequired("output"); err != nil {
		log.Fatal("failed to mark output as required")
	}
	if err := downloadIllustCmd.MarkFlagRequired("count"); err != nil {
		log.Fatal("failed to mark count as required")
	}
	illustCmd.AddCommand(downloadIllustCmd)
}
