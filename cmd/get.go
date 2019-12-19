/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"flag"
	"fmt"
	get "github.com/JoseThen/hn/util"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var Count int
var Top bool
var New bool

type Story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a list of HackerNews posts.",
	Long: `Retrieve a list of HackerNews posts.
	By default getting 10 Top posts.`,
	Run: func(cmd *cobra.Command, args []string) {
		w := new(tabwriter.Writer)
		// minwidth, tabwidth, padding, padchar, flags
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)
		count, _ := cmd.Flags().GetInt("count")
		top, _ := cmd.Flags().GetBool("top")
		new, _ := cmd.Flags().GetBool("new")

		flag.Parse()

		switch {
		case top:
			defer w.Flush()
			fmt.Fprintf(w, "\n %s\t%s\t", "Title", "Url")
			fmt.Fprintf(w, "\n %s\t%s\t", "-----", "---")
			var topIds = get.Ids(count, "top")
			for _, id := range topIds {
				var results = get.Data(id)
				fmt.Fprintf(w, "\n %s\t%s\t", results.Title, results.URL)
			}
		case new:
			defer w.Flush()
			fmt.Fprintf(w, "\n %s\t%s\t", "Title", "Url")
			fmt.Fprintf(w, "\n %s\t%s\t", "-----", "---")
			var newIds = get.Ids(count, "new")
			for _, id := range newIds {
				var results = get.Data(id)
				fmt.Fprintf(w, "\n %s\t%s\t", results.Title, results.URL)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().IntVarP(&Count, "count", "c", 10, "Number of posts to get.")
	getCmd.Flags().BoolVarP(&Top, "top", "t", true, "Get Top posts.")
	getCmd.Flags().BoolVarP(&New, "new", "n", false, "Get New posts.")
}
