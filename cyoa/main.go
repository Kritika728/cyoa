package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetStory)
	fmt.Println("Listen")
	http.ListenAndServe(":8080", nil)
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryIntro struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

func GetStory(w http.ResponseWriter, r *http.Request) {
	var data map[string]StoryIntro
	file, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatal("error while opening file: ", err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal("json unmarshal: ", err)
	}

	t, err := template.ParseFiles("story.html")
	if err != nil {
		log.Fatal("error while parshing html file: ", err)
	}

	err = t.Execute(w, data["intro"])
	if err != nil {
		log.Fatal("error here: ", err)
	}

}
