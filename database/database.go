package database

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/dmitryk-dk/blog/models"
	"github.com/dmitryk-dk/blog/config"
)

var dbInstance *sql.DB

func Connect (config *config.Config) (*sql.DB, error) {
	if dbInstance == nil {
		connectionConfig := fmt.Sprintf("%s:%s@%s/%s", config.User, config.Password, config.Host, config.DbName)
		db, err := sql.Open(config.DbDriverName, connectionConfig)
		if err != nil {
			return nil, err
		}
		dbInstance = db
	}
	return dbInstance, nil
}

type DbMethodsHelper interface {
	AddPost(*models.Post) error
	DeletePost(id int) error
	GetPost(*models.Post) error
	GetAllPosts() (models.Posts, error)
	//UpdatePost(*models.Post) error
}

type DbMethods struct{}

func (m *DbMethods) AddPost (post *models.Post) error {
	stmt, err := dbInstance.Prepare("INSERT post SET id=?,title=?,description=?")
	if err != nil {
		fmt.Errorf("Can't add to database: %s", err)
		return nil
	}
	res, err := stmt.Exec(post.Id, post.Title, post.Description)
	if err != nil {
		fmt.Errorf("Can't add to database: %s", err)
		return nil
	}
	fmt.Printf("%v\n", res)
	return nil
}

func (m *DbMethods) DeletePost(id int)  error {
	stmt, err := dbInstance.Prepare("DELETE from post where id=?")
	if err != nil {
		fmt.Errorf("Can't delete from database: %s", err)
		return nil
	}
	res, err := stmt.Exec(&id)
	if err != nil {
		fmt.Errorf("Can't add to database: %s", err)
		return nil
	}
	fmt.Printf("%v\n", res)
	return nil
}

func (m *DbMethods) GetPost (post *models.Post)  error {
	rows, err := dbInstance.Query("SELECT * FROM post")
	if err != nil {
		fmt.Errorf("Can't add to database: %s", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Description)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (m *DbMethods) GetAllPosts() (models.Posts, error) {
	posts := make(models.Posts, 0)
	rows, err := dbInstance.Query("SELECT * FROM post")
	if err != nil {
		fmt.Errorf("Can't add to database: %s", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		post := &models.Post{}
		rows.Scan(&post.Id, &post.Title, &post.Description)
		posts = append(posts, *post)
	}
	return posts, nil
}
//
//func (m *DbMethods) UpdatePost (post *models.Post) error {
//	return nil
//}
