package main

import (
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseFiles("static/home.html"))

	bio := "My Name is Lee Edbert Panti. I am 20 years old currently enrolled in the AINT program at the University of Belize. I love learning new things and figuring out solutions to programming related problems. This is my first Website using Golang programming language."
	html.Execute(w, bio)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseFiles("static/greeting.html"))

	dateTime := make(map[string]string)
	dateTime["time"] = time.Now().Format("15:04")
	dateTime["weekday"] = time.Now().Weekday().String()

	html.Execute(w, dateTime)
}

func random(w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseFiles("static/random.html"))

	quotes := [6]string{
		"Physics is the universe's operating system.",
		"A computer once beat me at chess, but it was no match for me at kick boxing.",
		"It's hardware that makes a machine fast.  It's software that makes a fast machine slow.",
		"There is only one problem with common sense; it's not very common.",
		"Programs must be written for people to read, and only incidentally for machines to execute.",
		"Before software should be reusable, it should be usable.",
	}

	randQuote := rand.Intn(len(quotes))

	html.Execute(w, quotes[randQuote])

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/greeting", greeting)
	mux.HandleFunc("/random", random)

	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles"))))

	log.Println("Starting server on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		panic(err)
	}
}
