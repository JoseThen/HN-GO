package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
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
	topCommand := flag.NewFlagSet("top", flag.ExitOnError)
	newCommand := flag.NewFlagSet("new", flag.ExitOnError)

	// Top subcommand flag pointers
	var topListNumber int
	topCommand.IntVar(&topListNumber, "num", 15, "Number of stories to show. Default is 15.")

	// New subcommand flag pointers
	// newListNumber := newCommand.Int("num", 15, "Number of stories to show. Default is 15.")

	if len(os.Args) < 2 {
		fmt.Println("You need to choose a subcommand")
		os.Exit(1)
	}

	/*
		Switch on the subcommand
		Parse the flags fore the appropriate FlagSet
		FlagSet.Parse() requires a set of arguments to parse as input
		os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	*/
	switch os.Args[1] {
	case "top":
		topCommand.Parse(os.Args[2:])
	case "new":
		newCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	if topCommand.Parsed() {
		defer w.Flush()
		fmt.Fprintf(w, "\n %s\t%s\t", "Title", "Url")
		fmt.Fprintf(w, "\n %s\t%s\t", "----", "----")
		var topIds = getIds(topListNumber, "top")
		for _, id := range topIds {
			var results = getData(id)
			fmt.Fprintf(w, "\n %s\t%s\t", results.Title, results.URL)
		}
	}

	if newCommand.Parsed() {
		defer w.Flush()
		fmt.Fprintf(w, "\n %s\t%s\t", "Title", "Url")
		fmt.Fprintf(w, "\n %s\t%s\t", "----", "----")
		var newIds = getIds(topListNumber, "new")
		for _, id := range newIds {
			var results = getData(id)
			fmt.Fprintf(w, "\n %s\t%s\t", results.Title, results.URL)
		}
	}
}

func getIds(amount int, category string) []int {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/%sstories.json", category)
	var topArray []int

	client := http.Client{
		Timeout: time.Second * 10, //Max of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal([]byte(body), &topArray)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer res.Body.Close()
	return topArray[:amount]

}

func getData(id int) Story {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	var story Story
	// fmt.Println(url)

	client := http.Client{
		Timeout: time.Second * 2, //Max of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal([]byte(body), &story)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// defer res.Body.Close()
	return story

}
