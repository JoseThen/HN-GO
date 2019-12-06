package get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Story hn Stories
type Story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func Ids(amount int, category string) []int {
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

func Data(id int) Story {
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
