package cmd

import (
	"fmt"
	"os"

	"github.com/cloud66-oss/trackman/utils"
	"github.com/khash/updater"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update alterant to the latest version",
	Run:   updateExec,
}

func init() {
	updateCmd.Flags().StringP("channel", "", utils.Channel, "version channel")
	_ = viper.BindPFlag("channel", updateCmd.Flags().Lookup("channel"))

	rootCmd.AddCommand(updateCmd)
}

func updateExec(cmd *cobra.Command, args []string) {
	update(false)

	fmt.Println("Updated")
}

func update(silent bool) {
	worker, err := updater.NewUpdater(utils.Version, &updater.Options{
		RemoteURL: "https://s3.amazonaws.com/downloads.cloud66.com/alterant/",
		Channel:   viper.GetString("channel"),
		Silent:    silent,
	})
	if err != nil {
		if !silent {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	err = worker.Run(viper.GetString("channel") != utils.Channel)
	if err != nil {
		if !silent {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Exit(0)
	}
}
