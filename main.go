package main

import (
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

// home page handler function
func home(w http.ResponseWriter, r *http.Request) {

	// creates a new template for the home page
	html := template.Must(template.ParseFiles("static/home.html"))

	// data to be added to the template
	bio := "My Name is Lee Edbert Panti. I am 20 years old currently enrolled in the AINT program at the University of Belize. I love learning new things and figuring out solutions to programming related problems. This is my first Website using Golang programming language."

	html.Execute(w, bio)
}

// greeting handler function
func greeting(w http.ResponseWriter, r *http.Request) {
	// template for the greeting page
	html := template.Must(template.ParseFiles("static/greeting.html"))

	// find the time and weekday to be displayed
	dateTime := make(map[string]string)
	dateTime["time"] = time.Now().Format("15:04")
	dateTime["weekday"] = time.Now().Weekday().String()

	html.Execute(w, dateTime)
}

// random handler function
func random(w http.ResponseWriter, r *http.Request) {
	// template for the random page
	html := template.Must(template.ParseFiles("static/random.html"))

	// quotes to be added to the template
	quotes := [6]string{
		"Physics is the universe's operating system.",
		"A computer once beat me at chess, but it was no match for me at kick boxing.",
		"It's hardware that makes a machine fast.  It's software that makes a fast machine slow.",
		"There is only one problem with common sense; it's not very common.",
		"Programs must be written for people to read, and only incidentally for machines to execute.",
		"Before software should be reusable, it should be usable.",
	}

	// passing in only one randum quote
	randQuote := rand.Intn(len(quotes))

	html.Execute(w, quotes[randQuote])

}

func main() {
	// create a new mux router
	mux := http.NewServeMux()
	// supply the end points and each handler function
	mux.HandleFunc("/", home)
	mux.HandleFunc("/greeting", greeting)
	mux.HandleFunc("/random", random)

	// supply the styles folder to let the templates use the stylesheet
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles"))))
	// /styles/styles.css

	// create a new server
	log.Println("Starting server on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		panic(err)
	}
}
