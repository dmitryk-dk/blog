package main

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/blog/models"
)



type PostJson struct {
	Data string
}

type ResponseOk struct {
	Ok 		bool `json:"success"`
}

type ResponseErr struct {
	Error	string	`json:"error"`
}

func indexHandler (w http.ResponseWriter, r *http.Request) {
	post := &models.Post{
		Id: 		 0,
		Title:   	 "New title",
		Description: "Description",
	}
	jsonPost, _ := json.Marshal(post)
	data := &PostJson{string(jsonPost) }
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", *data)
}

func dependenciesHandler () http.Handler {
	return http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/")))
}

func postHandler (w http.ResponseWriter, r *http.Request) {
	var post models.Post

	body,err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.Unmarshal(body, &post)

	posts := make(map[int]*models.Post, 0)
	posts[post.Id] = &post

	ok := &ResponseOk{ Ok: true }

	jsonResponse, _ := json.Marshal(ok)
	w.Write(jsonResponse)
}

func main () {
	port := "3030"
	depHandler := dependenciesHandler()
	http.Handle("/dist/", depHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/post", postHandler)

	fmt.Printf("Running server on port: %s\n Type Ctr-c to shutdown server.\n", port)

	http.ListenAndServe(":"+port, nil)
}
