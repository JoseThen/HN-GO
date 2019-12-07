package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	get "github.com/JoseThen/hn/util"
)

// Story hn Stories
type Story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	// SubCommands
	topFlag := flag.Bool("top", true, "Display the top HN posts.")
	newFlag := flag.Bool("new", true, "Display the newest HN posts.")
	countFlag := flag.Int("count", 10, "Number of posts to display.")

	flag.Parse()

	switch {
	case *topFlag:
		defer w.Flush()
		fmt.Fprintf(w, "\n %s\t%s\t", "Title", "Url")
		fmt.Fprintf(w, "\n %s\t%s\t", "----", "----")
		var topIds = get.Ids(*countFlag, "top")
		for _, id := range topIds {
			var results = get.Data(id)
			fmt.Fprintf(w, "\n %s\t%s\t", results.Title, results.URL)
		}
	case *newFlag:
		defer w.Flush()
		fmt.Fprintf(w, "\n %s\t%s\t", "Title", "Url")
		fmt.Fprintf(w, "\n %s\t%s\t", "----", "----")
		var newIds = get.Ids(*countFlag, "new")
		for _, id := range newIds {
			var results = get.Data(id)
			fmt.Fprintf(w, "\n %s\t%s\t", results.Title, results.URL)
		}
	}
}
