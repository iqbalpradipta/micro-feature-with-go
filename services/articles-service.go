package service

import (
	"github.com/iqbalpradipta/micro-feature-with-go/entities"
	"gorm.io/gorm"
)

type ArticleService interface {
	GetArticle() ([]entities.Articles, error)
	GetArticleById(id int) (entities.Articles, error)
}

func RepositoryArticle(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) GetArticle() ([]entities.Articles, error) {
	var articles []entities.Articles
	// var users []entities.Users
	err := r.db.Table("articles").Select("articles.title, articles.image, articles.description, articles.author, users.id").Joins("left join users on users.id = articles.id").Find(&articles).Error

	return articles, err
}

func (r *repository) GetArticleById(id int) (entities.Articles, error) {
	var articles entities.Articles
	err := r.db.Table("articles").Select("articles.title, articles.image, articles.description, articles.author, users.id").Joins("left join users on users.id = articles.id").First(&articles, id).Error

	return articles, err
}