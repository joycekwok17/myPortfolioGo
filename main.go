package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// 1) TODO: Fill out the personal_data.json file in the data directory

// Project struct to hold the data for each project
type Project struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Link        string `json:"Link"`
}

// 2) TODO: Create a Network struct to hold the data for each network
type Network struct {
	Name string `json:"Name"`
	Link string `json:"Link"`
}

// 3) TODO: Create a Portfolio struct to hold the data for the entire portfolio
// Hint: Notice that the attributs "Projects" and "Networks" must be slices/lists
type Portfolio struct {
	Name     string    `json:"Name"`
	Title    string    `json:"Title"`
	About    string    `json:"About"`
	Projects []Project `json:"Projects"`
	Networks []Network `json:"Networks"`
}

// 4) TODO: Go into templates/portfolio.gohtml and call the data from the Portfolio struct everywhere it says "TODO"

func main() {
	portfolio := handleData()
	createHTML(portfolio)
	startServer()
}

func handleData() Portfolio {
	// Read the data from the JSON file
	file, err := os.ReadFile("static/data/personal_data.json")
	if err != nil {
		log.Fatal("Failed to read personal_data.json:", err)
	}

	// Unmarshal the JSON data into a Portfolio struct
	var portfolio Portfolio
	err = json.Unmarshal(file, &portfolio)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON data:", err)
	}

	return portfolio
}

func createHTML(portfolio Portfolio) {
	// TODO: Create the static/index.html file
	file, err := os.Create("static/index.html")
	checkFileError(err)
	defer file.Close()
	// TODO: Parse the template file
	tmpl := template.Must(template.ParseFiles("templates/portfolio.gohtml"))

	// TODO: Execute the template and write the output to the index.html
	err = tmpl.Execute(file, portfolio)
}

func startServer() {
	// TODO: Create a file server to serve the files in the static directory
	fs := http.FileServer(http.Dir("."))

	// TODO: Handle the root path "/"
	http.Handle("/", fs)

	// TODO: Start the server, listen on port 8080
	fmt.Println("server starts listening on port 8090")
	err := http.ListenAndServe(":8090", nil)
	checkFileError(err)
}

func checkFileError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
