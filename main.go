package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("go-htmx started! ", time.Now().Format("2006-01-02 15:04:05"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{"Everything Everywhere All at Once", "Daniel Kwan, Daniel Scheinert"},
				{"The Matrix", "Lana Wachowski, Lilly Wachowski"},
				{"The Shawshank Redemption", "Frank Darabont"},
				{"Wonder Woman", "Patty Jenkins"},
				{"Interstellar", "Christopher Nolan"},
				{"Anatomy of a Fall", "Justine Triet"},
			},
		}

		tmpl.Execute(w, films)
	})

	http.HandleFunc("/add-film/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMLX request received")
		log.Print(r.Header.Get("HX-Request"))
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
