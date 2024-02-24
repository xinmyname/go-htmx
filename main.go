package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("go-htmx started! ", time.Now().Format("2006-01-02 15:04:05"))

	mux := http.NewServeMux()

	fileServer := http.FileServer(safeStaticFileSystem{http.Dir("./static")})

	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("/add-film/", func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		time.Sleep(2 * time.Second)

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}

type safeStaticFileSystem struct {
	fs http.FileSystem
}

func (ssfs safeStaticFileSystem) Open(path string) (http.File, error) {
	f, err := ssfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, _ := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := ssfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
