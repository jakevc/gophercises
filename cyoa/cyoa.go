// Choose your own adventure (cyoa), is an interactive story structured as a webapp.
package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"html/template"
)

// Cyoa struct holds stories and their contents.
type Cyoa struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
}

// Story is map that contains each chapter's story as a Cyoa struct.
type Story map[string]Cyoa

// parseJSON parses json from gopher.json into a map of chapters.
func parseJSON() map[string]Cyoa {

	// story data
	var adventure Story

	// read in the gopher.json story data
	file, err := os.Open("gopher.json")
	if err != nil {
		log.Fatalf("Failed to open json file: %s", err)
	}
	defer file.Close() 

	// fil to bytes
	cyoaData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to convert chapters to bytes: %s", err)
	}

	// unmarshal json into Story struct
	err = json.Unmarshal(cyoaData, &adventure) 
	if err != nil {
		log.Fatal(err)
	}

	return adventure
}

// parse each story into the template
func storyHandler(adventureMap map[string]Cyoa) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arc := r.URL.Path[len("/arc/"):]
		if _ , ok := adventureMap[arc]; ok {
			t := template.Must(template.ParseFiles("tmpl/story.html"))
			err := t.Execute(w, adventureMap[arc])
			if err != nil {
				log.Fatal(err)	
			}
		}

	}
}

// the index handler shows the home page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	ind := template.Must(template.ParseFiles("tmpl/index.html"))
	if err := ind.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}


func main() {
	a := parseJSON()
	Story := storyHandler(a)
	fmt.Println("Serving Cyoa on: http://localhost:8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/arc/", Story)
	http.ListenAndServe(":8080", nil)
}


