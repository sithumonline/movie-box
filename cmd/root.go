package cmd

import (
	"os"

	"github.com/sithumonline/movie-box/cmd/get"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any sub commands
var rootCmd = &cobra.Command{
	Use:   "movie-box",
	Short: "Download Movies",
	Long:  `Download movies by name from YTS without visiting to YTS`,
}

func init() {
	rootCmd.AddCommand(get.GetMovieCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorf(err.Error())
		os.Exit(1)
	}
}
