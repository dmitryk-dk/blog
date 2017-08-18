package main

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/dmitryk-dk/blog/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

var posts map[int]*models.Post

func indexHandler (w http.ResponseWriter, r *http.Request) {
	post := &models.Post{
		Id: 		 0,
		Title:   	 "New title",
		Description: "Description",
	}
	jsonPost, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data := &PostJson{string(jsonPost) }
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "index", *data)
}

func dependenciesHandler () http.Handler {
	return http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/")))
}

func postHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var post models.Post
		body,err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = json.Unmarshal(body, &post)
		if err != nil {
			errorResp := &ResponseErr{"You don't create post"}
			jsonErrResponse, err := json.Marshal(errorResp)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Write(jsonErrResponse)
		}
		posts[post.Id] = &post
		ok := &ResponseOk{ Ok: true }
		jsonResponse, err := json.Marshal(ok)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(jsonResponse)
	} else {
		errorMethod := &ResponseErr{"You use error method"}
		jsonErrResponse, _ := json.Marshal(errorMethod)
		w.Write(jsonErrResponse)
		http.Error(w, "Used another Method", http.StatusInternalServerError)
	}
}

func deleteHandler (w http.ResponseWriter, r *http.Request) {
	var id int

	body,err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.Unmarshal(body, &id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	delete(posts, id)
	ok := &ResponseOk{ Ok: true }
	jsonResponse, err := json.Marshal(ok)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonResponse)
}

func main () {
	const port = "3030"
	posts = make(map[int]*models.Post, 0)
	db, err := sql.Open("mysql", "dmitryk:DbImOK85!@tcp(192.168.56.107:3306)/posts")
	depHandler := dependenciesHandler()
	http.Handle("/dist/", depHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/delete", deleteHandler)
	if err != nil {
		panic(err)
	}
	res, err := db.Query("SELECT * FROM `post`")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("%v\n", res)
	db.Close()
	fmt.Printf("Running server on port: %s\n Type Ctr-c to shutdown server.\n", port)

	http.ListenAndServe(":"+port, nil)
}
