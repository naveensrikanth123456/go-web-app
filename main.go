package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Home page accessed")
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Courses page accessed")
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	log.Println("About page accessed")
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Contact page accessed")
	http.ServeFile(w, r, "static/contact.html")
}

func main() {

	log.Println("Go Web App Started")
	log.Println("Listening on port 8080")

	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
