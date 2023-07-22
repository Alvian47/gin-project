package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type Article struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rate        uint   `json:"rate"`
}

func getEnvVariable(key string) string {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func NewPostgreSQLClient() {
	var (
		host     = getEnvVariable("DB_HOST")
		port     = getEnvVariable("DB_PORT")
		user     = getEnvVariable("DB_USER")
		dbname   = getEnvVariable("DB_NAME")
		password = getEnvVariable("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(Article{})
}

func CreateArticle(a *Article) (*Article, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &Article{}, errors.New("article not created")
	}

	return a, nil
}

func ReadArticle(id string) (*Article, error) {
	var article Article
	res := db.First(&article, id)
	if res.RowsAffected == 0 {
		return &Article{}, errors.New("article not created")
	}
	return &article, nil
}

func ReadArticles() ([]*Article, error) {
	var articles []*Article
	res := db.Find(&articles)
	if res.Error != nil {
		return nil, errors.New("authors not created")
	}

	return articles, nil
}

func UpdateArticle(article *Article) (*Article, error) {
	var updateArticle Article
	result := db.Model(&updateArticle).Where(article.ID).Updates(article)
	if result.Error != nil {
		return &Article{}, result.Error
	}
	if result.RowsAffected == 0 {
		return &Article{}, fmt.Errorf("id tidak ada")
	}
	return &updateArticle, nil
}

func DeleteArticle(id string) error {
	var deleteArticle Article
	result := db.Where(id).Delete(&deleteArticle)
	if result.RowsAffected == 0 {
		return errors.New("article not created")
	}
	return nil
}
