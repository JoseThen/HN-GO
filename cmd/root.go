package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "hn",
  Short: "CLI to get HAcker News Posts and Information",
	Long: `hn is a cli utility to get information on the
	current posts in Hacker News. It takes advantage of the
	Hacker News API.`
}

// // Execute executes the root command.
// func Execute() error {
// 	return rootCmd.Execute()
// }

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
}
