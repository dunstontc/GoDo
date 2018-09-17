package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var theTodos Todos

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func getTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		desc := r.Form["description"][0]

		if len(desc) > 0 {
			theTodos = append(theTodos, Todo{desc, getTimestamp(), false})
		}
	}

	http.Redirect(w, r, "/", 302)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		date := r.Form["date"][0]

		theTodos = Drop(date, theTodos)
	}
	http.Redirect(w, r, "/", 302)
}

func toggleTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		date := r.Form["date"][0]
		index := getIndex(date, theTodos)

		theTodos[index].toggle()
	}
	http.Redirect(w, r, "/", 302)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		date := r.Form["date"][0]
		description := r.Form["description"][0]

		index := getIndex(date, theTodos)
		theTodos[index].update(description)

		http.Redirect(w, r, "/", 302)

	} else {
		dat := struct {
			Todos []Todo
			Year  int
		}{
			theTodos,
			time.Now().Year(),
		}

		log.Println(theTodos)
		tpl.ExecuteTemplate(w, "update.gohtml", dat)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	dat := struct {
		Todos []Todo
		Year  int
	}{
		theTodos,
		time.Now().Year(),
	}

	log.Println(theTodos)
	tpl.ExecuteTemplate(w, "index.gohtml", dat)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/go.ico")
}

func init() {
	// tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.HandleFunc("/favicon.ico", faviconHandler)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/todos/new", createTodo)
	http.HandleFunc("/todos/toggle", toggleTodo)
	http.HandleFunc("/todos/update", updateTodo)
	http.HandleFunc("/todos/delete", deleteTodo)

	log.Println("Listening on localhost: " + getPort())
	err := http.ListenAndServe(getPort(), nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}