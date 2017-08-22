package main

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"os/signal"
	"os"
	"syscall"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/dmitryk-dk/blog/models"
	"github.com/dmitryk-dk/blog/config"
	"github.com/dmitryk-dk/blog/database"
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
	var dbHelper database.DbMethodsHelper
	dbHelper = &database.DbMethods{}
	posts, err := dbHelper.GetAllPosts()
	jsonPost, err := json.Marshal(posts)
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
	var dbHelper database.DbMethodsHelper
	var post *models.Post
	if r.Method == "POST" {
		post = &models.Post{}
		dbHelper = &database.DbMethods{}
		body,err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err = json.Unmarshal(body, post); err != nil {
			log.Println("Unmarshall error:", err)
		}

		err = dbHelper.AddPost(post)
		if err != nil {
			errorResp := &ResponseErr{"You don't create post"}
			jsonErrResponse, err := json.Marshal(errorResp)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Write(jsonErrResponse)
		}

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
	var dbHelper database.DbMethodsHelper
	if r.Method == "DELETE" {
		dbHelper = &database.DbMethods{}
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		json.Unmarshal(body, &id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = dbHelper.DeletePost(id)

		if err != nil {
			errorResp := &ResponseErr{"You don't create post"}
			jsonErrResponse, err := json.Marshal(errorResp)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Write(jsonErrResponse)
		}

		ok := &ResponseOk{Ok: true }
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

func main () {
	const port = "3030"
	cfg := config.GetConfig()
	db, err := database.Connect(cfg)
	if err != nil {
		panic(err)
	}
	depHandler := dependenciesHandler()
	http.Handle("/dist/", depHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/delete", deleteHandler)
	preparedShutdown(db)
	fmt.Printf("Running server on port: %s\n Type Ctr-c to shutdown server.\n", port)
	http.ListenAndServe(":"+port, nil)
}

func preparedShutdown(db *sql.DB) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Got signal: %d", <-sig)
		db.Close()
		os.Exit(0)
	}()
}
